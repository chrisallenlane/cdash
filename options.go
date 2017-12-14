package main

import (
	"os"
)

type Options struct {
	ConfigFile string
}

func NewOptions(docopts map[string]interface{}) Options {

	// initialize an Options object
	options := Options{}

	// ConfigFile
	if docopts["--config"] != nil {
		options.ConfigFile = docopts["--config"].(string)
	} else {
		options.ConfigFile = os.Getenv("HOME") + "/.config/cdash.yml"
	}

	// return the options object
	return options
}
