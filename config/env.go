package config

import (
	"os"
	"strconv"
)

// NOTE: "direnv" package is used for loading environment variables

type EnvConfig struct {
	Port int
}

func InitEnvConfig() EnvConfig {
	return EnvConfig{
		Port: getInt("PORT", 8080),
	}
}

func getInt(key string, fallback int) int {
	valStr, set := os.LookupEnv(key)
	if !set {
		return fallback
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}

	return val
}
