package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"os"
)

func main() {
	// initialize options
	docopts, _ := docopt.Parse(usage(), nil, true, "1.1.3", false)
	options := newOptions(docopts)

	// initialize configs
	config, err := newConfig(options)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetFlags(0) // disable timestamps

	// query the coinmarketcap API
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=0")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalln(fmt.Sprintf(
			"CoinMarketCap API responded with HTTP status %d.",
			response.StatusCode,
		))
	}

	// read the response body and unpack the JSON response
	var coins []coin
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatalln(err)
	}

	// restructure coin data into a map
	hash := map[string]coin{}
	for _, coin := range coins {
		hash[coin.Symbol] = coin
	}

	// assemble the coin portfolio
	portfolio := []coin{}
	for _, coin := range config {
		// throw an error if no CoinMarketCap data exists for `coin`
		if _, exists := hash[coin.Symbol]; exists == false {
			log.Fatalln(errors.New("No data available for token \"" + coin.Symbol + "\"."))
		}

		// merge coin data
		c := hash[coin.Symbol]
		c.Holdings = coin.Holdings
		c.Worth = coin.Holdings

		// add coin to portfolio
		portfolio = append(portfolio, c)
	}

	// aggregate table data
	var total float64
	rows := make([][]string, 1)
	for _, coin := range portfolio {
		worth := coin.Holdings * coin.PriceUSD
		total += worth
		row := []string{
			coin.Symbol,
			format("$", coin.PriceUSD, false),
			format("%", coin.Delta1H, true),
			format("%", coin.Delta1D, true),
			format("%", coin.Delta7D, true),
			format("", coin.Holdings, false),
			format("$", worth, false),
		}
		rows = append(rows, row)
	}

	// display table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Token",
		"USD",
		"Δ 1H",
		"Δ 24H",
		"Δ 7D",
		"Holdings",
		"Worth",
	})

	// append table data
	table.AppendBulk(rows)

	// table footer
	table.SetFooter([]string{
		"",
		"",
		"",
		"",
		"",
		"Portfolio Value",
		format("$", total, false),
	})

	// set column alignment
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
	})

	// write to stdout
	table.Render()
}
