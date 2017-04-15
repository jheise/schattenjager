package main

import (

	// standard
	"os"

	// external
	"github.com/burntsushi/toml"
)

type MainConfig struct {
	Verbose   bool
	Publisher string
	PubConfig map[string]interface{}
}

func ReadConfig(configsrc string) (*MainConfig, error) {
	var config MainConfig

	_, err := os.Stat(configsrc)
	if err != nil {
		return &config, err
	}

	_, err = toml.DecodeFile(configsrc, &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}
