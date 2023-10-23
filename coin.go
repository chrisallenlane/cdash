package main

import (
	"errors"
	"strconv"
)

type coin struct {
	ID       string
	Name     string
	Symbol   string
	Price    float64
	Delta1H  float64
	Delta1D  float64
	Delta7D  float64
	Holdings float64
	Worth    float64
}

// constructs a coin from raw JSON
func newCoin(cmc CoinMarketCoin, base string) coin {
	var d1h float64
	var d1d float64
	var d7d float64


	d1h = cmc.Quote[base].Percent_change_1h
	d1d = cmc.Quote[base].Percent_change_24h
	d7d = cmc.Quote[base].Percent_change_7d

	// parse out the price in the specified currency
	// initialize and return the coin
	return coin{
		ID:      strconv.Itoa(cmc.ID),
		Name:    cmc.Name,
		Symbol:  cmc.Symbol,
		Delta1H: d1h,
		Delta1D: d1d,
		Delta7D: d7d,
		Price:   cmc.Quote[base].Price,
	}
}

// returns the currency symbol associated with the currency code
func cSym(sym string) (string, error) {

	// map currency codes to symbols
	symbols := map[string]string{
		"AUD": "$",
		"BRL": "R$",
		"CAD": "$",
		"CLP": "$",
		"CNY": "¥",
		"CZK": "Kč",
		"DKK": "kr",
		"EUR": "€",
		"HKD": "$",
		"HUF": "Ft",
		"INR": "₹",
		"IDR": "Rp",
		"ILS": "₪",
		"JPY": "¥",
		"KRW": "₩",
		"MYR": "RM",
		"MXN": "$",
		"NZD": "$",
		"NOK": "kr",
		"PKR": "₨",
		"PHP": "₱",
		"PLN": "zł",
		"RUB": "₽",
		"SGD": "$",
		"ZAR": "R",
		"SEK": "kr",
		"CHF": "Fr",
		"TWD": "NT$",
		"THB": "฿",
		"TRY": "₺",
		"GBP": "£",
		"USD": "$",
	}

	// return an error if the specified currency is not supported
	if _, ok := symbols[sym]; ok == false {
		return "", errors.New("Currency " + sym + " is not supported.")
	}

	// return the currency symbol
	return symbols[sym], nil
}
