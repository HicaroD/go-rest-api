package envloader

import (
	"os"
	"strconv"
)

// TODO: use package "reflect" for reading field tags from
// struct and retrieve the environment variables
//
// The idea is to be really simple because environment
// variables are generally strings or integers

func GetInt(key string, fallback int) int {
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

func GetString(key string, fallback string) string {
	val, set := os.LookupEnv(key)
	if !set {
		return fallback
	}
	return val
}
