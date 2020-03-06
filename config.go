package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type CdashConfig struct {
	Coins []coin
	Key string
}

// method that returns config struct
func newConfig(opts options) (CdashConfig, error) {

	// read the config file
	buf, err := ioutil.ReadFile(opts.ConfigFile)
	if err != nil {
		return CdashConfig{}, err
	}

	//initialize a config object
	var config CdashConfig

	// unmarshal the yaml
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return CdashConfig{}, err
	}

	// return the configs
	return config, nil
}
