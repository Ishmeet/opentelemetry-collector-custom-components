package sflowreceiver

import (
	"go.opentelemetry.io/collector/config/confignet"
)

type Config struct {
	confignet.NetAddr `mapstructure:",squash"`
	Labels            map[string]string `mapstructure:"labels"`
}

func (cfg *Config) Validate() error {
	return nil
}
