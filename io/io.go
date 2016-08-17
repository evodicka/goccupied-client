package io

import (
	"github.com/BurntSushi/toml"
)

const resourcePath string = "./"

func ReadConfig() *config {
	var appConfig config

	toml.DecodeFile(resourcePath+"config.toml", &appConfig)
	return &appConfig
}
