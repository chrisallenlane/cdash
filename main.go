package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chrisallenlane/thou"
	"github.com/docopt/docopt-go"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// disable logger timestamps
	log.SetFlags(0)

	// initialize options
	docopts, _ := docopt.Parse(usage(), nil, true, "1.2.1", false)
	options, err := newOptions(docopts)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize configs
	config, err := newConfig(options)
	if err != nil {
		log.Fatalln(err)
	}

	// assemble the API URL
	url := strings.Join([]string{
		"https://api.coinmarketcap.com/v1/ticker/",
		"?convert=",
		options.Base,
		"&limit=0",
	}, "")

	// query the coinmarketcap API
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// assert a 200 HTTP response
	if response.StatusCode != http.StatusOK {
		log.Fatalln(fmt.Sprintf(
			"CoinMarketCap API responded with HTTP status %d.",
			response.StatusCode,
		))
	}

	// read and unpack the resposne body into a "raw" JSON object
	var raw []map[string]string
	if err := json.NewDecoder(response.Body).Decode(&raw); err != nil {
		log.Fatalln(err)
	}

	// iterate over the raw JSON elements and create a map of symbols to coins
	hash := map[string]coin{}
	for _, item := range raw {
		coin := newCoin(item, options.Base)
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

		// tally total portfolio value
		worth := coin.Holdings * coin.Price
		total += worth

		// apply thousands formatting
		fPrice, _ := thou.SepF(coin.Price, 2, ",", ".")
		fD1h, _ := thou.SepF(coin.Delta1H, 2, ",", ".")
		fD1d, _ := thou.SepF(coin.Delta1D, 2, ",", ".")
		fD7d, _ := thou.SepF(coin.Delta7D, 2, ",", ".")
		fHoldings, _ := thou.SepF(coin.Holdings, 8, ",", ".")
		fWorth, _ := thou.SepF(worth, 2, ",", ".")

		row := []string{
			coin.Symbol,
			options.Symbol + fPrice,
			colorize(fD1h+"%", coin.Delta1H),
			colorize(fD1d+"%", coin.Delta1D),
			colorize(fD7d+"%", coin.Delta7D),
			fHoldings,
			options.Symbol + fWorth,
		}
		rows = append(rows, row)
	}

	// format the total
	fTotal, _ := thou.SepF(total, 2, ",", ".")

	// display table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Token",
		options.Base,
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
		options.Symbol + fTotal,
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
