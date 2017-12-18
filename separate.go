package main

import (
	"strconv"
	"strings"
)

func separate(value float64, precision int, thou string, dec string) string {
	// convert the float to a string
	str := strconv.FormatFloat(value, 'f', precision, 64)

	// create a buffer for the result
	formatted := ""

	// split the string on the integers/decimals separator
	parts := strings.Split(str, dec)
	ints := parts[0]
	decs := parts[1]

	// iterate backwards across the integers
	for i, j := len(ints)-1, 0; i >= 0; i-- {
		val := string(ints[i])

		// append a thousands separator after every 3rd digit
		if j%3 == 0 {
			val += thou
		}
		formatted = val + formatted
		j++
	}

	// trim the unwanted separator, concat the decimals, and return
	return strings.Trim(formatted, thou) + dec + decs
}
