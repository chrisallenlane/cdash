package main

import (
	"errors"
	"strconv"
	"strings"
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
func newCoin(json map[string]string, base string) coin {
	var d1h float64
	var d1d float64
	var d7d float64

	if val, ok := json["percent_change_1h"]; ok {
		d1h, _ = strconv.ParseFloat(val, 10)
	}

	if val, ok := json["percent_change_24h"]; ok {
		d1d, _ = strconv.ParseFloat(val, 10)
	}

	if val, ok := json["percent_change_7d"]; ok {
		d7d, _ = strconv.ParseFloat(val, 10)
	}

	// parse out the price in the specified currency
	b := strings.ToLower(base)
	price, _ := strconv.ParseFloat(json["price_"+b], 10)

	// initialize and return the coin
	return coin{
		ID:      json["id"],
		Name:    json["name"],
		Symbol:  json["symbol"],
		Delta1H: d1h,
		Delta1D: d1d,
		Delta7D: d7d,
		Price:   price,
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
