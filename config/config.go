package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config defines the properties of environment key
type Config struct {
	Driver  string `envconfig:"DRIVER"`
	Sqlconn string `envconfig:"SQLCONN"`
}

// FromEnv get all environment keys
func FromEnv() (*Config, error) {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		return &config, err
	}
	return &config, nil
}
