package main

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/mgutz/ansi"
)

// colorizes a string
func colorize(str string, val float64) string {
	// helpers
	green := ansi.ColorFunc("green")
	red := ansi.ColorFunc("red")

	// don't colorize if output is not being written to a tty
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}

	// colorize per val
	switch {
	case val > 0:
		str = green(str)
	case val < 0:
		str = red(str)
	}

	return str
}
