package main

import (
	"os"
	"strings"
)

type options struct {
	ConfigFile string
	Base       string
	Symbol     string
}

func newOptions(docopts map[string]interface{}) (options, error) {

	// ConfigFile
	var configFile string
	if docopts["--config"] != nil {
		configFile = docopts["--config"].(string)
	} else {
		configFile = os.Getenv("HOME") + "/.config/cdash.yml"
	}

	// Base currency
	base := strings.ToUpper(docopts["--base"].(string))

	// Currency symbol
	symbol, err := cSym(base)
	if err != nil {
		return options{}, err
	}

	// initialize an options object
	opts := options{
		ConfigFile: configFile,
		Base:       base,
		Symbol:     symbol,
	}

	// return the opts object
	return opts, nil
}
