package main

import (
	"encoding/json"
	"errors"
	"github.com/docopt/docopt-go"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// initialize options
	docopts, _ := docopt.Parse(usage(), nil, true, "1.0.0", false)
	options := NewOptions(docopts)

	// initialize configs
	config, err := NewConfig(options)
	if err != nil {
		log.Fatalln(err)
	}

	// query the coinmarketcap API
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// unpack the JSON response
	coins := []Coin{}
	err = json.Unmarshal([]byte(body), &coins)
	if err != nil {
		log.Fatalln(err)
	}

	// restructure coin data into a map
	hash := map[string]Coin{}
	for _, coin := range coins {
		hash[coin.Symbol] = coin
	}

	// assemble the coin portfolio
	portfolio := []Coin{}
	for _, coin := range config {

		// throw an error if no CoinMarketCap data exists for `coin`
		if _, exists := hash[coin.Symbol]; exists == false {
			log.Fatalln(errors.New("No data available for token '" + coin.Symbol + "'."))
		}

		// merge coin data
		c := hash[coin.Symbol]
		c.Holdings = coin.Holdings
		c.Worth = coin.Holdings

		// add coin to portfolio
		portfolio = append(portfolio, c)
	}

	// aggregate table data
	var total float64 = 0
	rows := make([][]string, 1)
	for _, coin := range portfolio {
		worth := coin.Holdings * coin.PriceUSD
		total += worth
		row := []string{
			coin.Symbol,
			format(coin.PriceUSD, "", false),
			format(coin.Delta1H, "%", true),
			format(coin.Delta1D, "%", true),
			format(coin.Delta7D, "%", true),
			format(coin.Holdings, "", true),
			format(worth, "", false),
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
		"$" + strconv.FormatFloat(total, 'f', 2, 64),
	})

	// write to stdout
	table.Render()
}
