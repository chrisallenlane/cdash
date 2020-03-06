package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chrisallenlane/thou"
	"github.com/docopt/docopt-go"
	"github.com/olekukonko/tablewriter"
)

type CoinMarketStatus struct {
	timestamp string
	error_code int
	error_message string
	elapsed int
	credit_count int
	notice string
}

type CoinMarketQuote struct {
	Price float64
	Volume_24h float64
	Percent_change_1h float64
	Percent_change_24h float64
	Percent_change_7d float64
	Market_cap float64
	Last_updated string
}

type CoinMarketCoin struct {
	ID int
	Name string
	Symbol string
	Slug string
	Date_added string
	Max_supply float64
	Circulating_supply float64
	Total_supply float64
	Last_updated string
	Quote map[string]CoinMarketQuote
}

type CoinMarketRes struct {
	Status CoinMarketStatus `json:"status"`
	Data []CoinMarketCoin `json:"data"`
}

func main() {
	// disable logger timestamps
	log.SetFlags(0)

	// initialize options
	docopts, _ := docopt.Parse(usage(), nil, true, "1.3.0", false)
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
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest",
		"?convert=",
		options.Base,
		"&start=1",
		"&limit=5000",
	}, "")

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", config.Key )

	// query the coinmarketcap API
	response, err := client.Do(req)
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

	// read and unpack the response body into a "raw" JSON object
	var raw CoinMarketRes
	if err := json.NewDecoder(response.Body).Decode(&raw); err != nil {
		log.Fatalln(err)
	}

	// iterate over the raw JSON elements and create lookups by Symbol and ID
	bySym := map[string][]coin{}
	byID := map[string]coin{}
	for _, cmc := range raw.Data {
		coin := newCoin(cmc, options.Base)
		bySym[coin.Symbol] = append(bySym[coin.Symbol], coin)
		byID[coin.ID] = coin
	}

	// assemble the coin portfolio
	portfolio := []coin{}
	for _, cn := range config.Coins {

		var c coin

		// look up by symbol
		//
		// NB: coin symbols are NOT unique. (See issue #4). As such, test to see if
		// the user has specified an ambiguous symbol in the configs. If an
		// ambiguous symbol is found, prompt the user to specify an ID instead.
		if e, exists := bySym[cn.Symbol]; exists {

			// if elements share a common symbol
			if len(bySym[cn.Symbol]) != 1 {

				// assemble an error message
				ids := ""
				for _, cx := range e {
					ids += fmt.Sprintf("\"%s\",", cx.ID)
				}
				ids = strings.Trim(ids, ",")

				// display it, and terminate
				log.Fatalln(fmt.Errorf(
					"Symbol \"%s\" is ambiguous. Please specify an ID instead. (One of: %s)",
					cn.Symbol,
					ids,
				))
			}

			// otherwise, take the first element if the symbol is unique
			c = e[0]

			// look up by ID
		} else if e, exists := byID[cn.ID]; exists {
			c = e

			// throw an error if no CoinMarketCap data exists for `cn`
		} else {
			log.Fatalln(fmt.Errorf("No data available for \"%s\".", cn.Symbol))
		}

		// merge coin data
		c.Holdings = cn.Holdings
		c.Worth = cn.Holdings

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

		// handle --name
		token := coin.Symbol
		if options.Name == true {
			token = coin.Symbol + " - " + coin.Name
		}

		row := []string{
			token,
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
