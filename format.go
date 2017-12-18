package main

import (
	"github.com/mattn/go-isatty"
	"github.com/mgutz/ansi"
	"os"
)

func format(what string, value float64, colorize bool) string {

	// colorizing helpers
	green := ansi.ColorFunc("green")
	red := ansi.ColorFunc("red")

	// format the value
	formatted := ""
	switch what {
	case "%":
		formatted = separate(value, 2, ",", ".") + "%"
	case "":
		formatted = separate(value, 8, ",", ".")
	case "$":
		formatted = "$" + separate(value, 2, ",", ".")
	}

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
