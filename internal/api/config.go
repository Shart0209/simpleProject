package api

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"time"
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
	Host           string        `envconfig:"POSTGRES_HOST" default:"localhost" required:"true"`
	Port           string        `envconfig:"POSTGRES_PORT" default:"5432" required:"true"`
	User           string        `envconfig:"POSTGRES_USER" default:"postgres" required:"true"`
	Pswd           string        `envconfig:"POSTGRES_PASSWORD" default:"1234qwER" required:"true"`
	DbName         string        `envconfig:"POSTGRES_DB_NAME" default:"test_db" required:"true"`
	ConnectTimeout time.Duration `envconfig:"POSTGRES_CONNECT_TIMEOUT" default:"1m" required:"true"`
	MaxConns       int           `envconfig:"POSTGRES_MAX_CONNS" default:"32"`
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
