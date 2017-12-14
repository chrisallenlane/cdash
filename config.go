package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// method that returns config struct
func NewConfig(options Options) ([]Coin, error) {

	// read the config file
	buf, err := ioutil.ReadFile(options.ConfigFile)
	if err != nil {
		return nil, err
	}

	//initialize a config object
	config := make([]Coin, 1)

	// unmarshal the yaml
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	// return the configs
	return config, nil
}
