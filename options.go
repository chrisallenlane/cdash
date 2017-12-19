package main

import (
	"os"
)

type options struct {
	ConfigFile string
}

func newOptions(docopts map[string]interface{}) options {

	// initialize an options object
	opts := options{}

	// ConfigFile
	if docopts["--config"] != nil {
		opts.ConfigFile = docopts["--config"].(string)
	} else {
		opts.ConfigFile = os.Getenv("HOME") + "/.config/cdash.yml"
	}

	// return the opts object
	return opts
}
