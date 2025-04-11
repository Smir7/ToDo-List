package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT" default:"8080"`
}

var AppConfig Config

func LoadConfig() error {
	return envconfig.Process("", &AppConfig)
}
