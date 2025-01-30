package config

import "lego-api-go/pkg/envloader"

// NOTE: "direnv" package is used for loading environment variables

type EnvConfig struct {
	Port int
}

func InitEnvConfig() EnvConfig {
	return EnvConfig{
		Port: envloader.GetInt("PORT", 8080),
	}
}
