package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is the configuration for the service
type Config struct {
	Database Database `envconfig:"database"`
}

type Database struct {
	DriverName string `envconfig:"driver_name" default:"postgres"`
	// TODO: Make this use SSL
	DataSourceName string `envconfig:"data_source_name" default:"postgres://postgres:postgres@localhost/service?sslmode=disable"`
}

// New returns a new Config struct
func New() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
