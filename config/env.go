package config

import "lego-api-go/pkg/envloader"

// NOTE: "direnv" package is used for loading environment variables
// Make sure "direnv" in your machine

type EnvConfig struct {
	Port int `env:"PORT" required:"true" default:"8080"`
}

func InitEnvConfig() *EnvConfig {
	cfg := EnvConfig{}
	err := envloader.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
