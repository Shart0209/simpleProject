package api

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	IsDebug       string `envconfig:"IS_DEBUG" default:"true"`
	LogLevel      string `envconfig:"LOG_LEVEL" default:"debug"`
	Listen        string `envconfig:"LISTEN" default:":8080"`
	FilesFolder   string `envconfig:"FILES_FOLDER" default:"upload"`
	Postgres      *PostgresConfig
	Authorization *AuthorizationConfig
}

type PostgresConfig struct {
	PostgresPSWD string `envconfig:"POSTGRES_PASSWORD" default:"1234qwER"`
	PostgresUSER string `envconfig:"POSTGRES_USER" default:"root"`
	PostgresDB   string `envconfig:"POSTGRES_DB" default:"postgresDB"`
}

type AuthorizationConfig struct {
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Error().Err(err).Msg("failed reading config")
		return nil, err
	}
	return &config, nil
}
