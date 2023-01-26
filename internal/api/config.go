package api

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	Listen   string `envconfig:"LISTEN" default:":8888"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Error().Err(err).Msg("failed reading config")
		return nil, err
	}
	return &config, nil
}
