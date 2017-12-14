package main

import (
	"github.com/mattn/go-isatty"
	"github.com/mgutz/ansi"
	"os"
	"strconv"
)

func format(value float64, unit string, colorize bool) string {

	// colorizing helpers
	green := ansi.ColorFunc("green")
	red := ansi.ColorFunc("red")

	// format the value
	formatted := strconv.FormatFloat(value, 'f', 2, 64) + unit

	// return early if no colorization is to be performed
	if !isatty.IsTerminal(os.Stdout.Fd()) || colorize == false {
		return formatted
	}

	// otherwise, colorize
	switch {
	case value > 0:
		formatted = green(formatted)
	case value < 0:
		formatted = red(formatted)
	}

	// return the formatted value
	return formatted
}
