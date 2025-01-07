package config

import "github.com/HicaroD/api/pkg"

// NOTE: "direnv" package is used for loading environment variables

type EnvConfig struct {
	Port int
}

func InitEnvConfig() EnvConfig {
	return EnvConfig{
		Port: utils.EnvGetInt("PORT", 8080),
	}
}
