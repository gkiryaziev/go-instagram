package conf

import (
	"github.com/jinzhu/configor"
)

var (
	config *Config
)

func GetSettings() *Config {
	if config == nil {
		configor.Load(&config, "config.yaml")
	}
	return config
}

type Config struct {
	Main struct {
		Instagram Instagram
	}
}

type Instagram struct {
	Login    string
	Password string
}
