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
	Postgres      *StorageConfig
	Authorization *AuthorizationConfig
}

type StorageConfig struct {
	Host string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port string `envconfig:"POSTGRES_PORT" default:"5432"`
	User string `envconfig:"POSTGRES_USER" default:"root"`
	Pswd string `envconfig:"POSTGRES_PASSWORD" default:"1234qwER"`
	Db   string `envconfig:"POSTGRES_DB" default:"postgres"`
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
