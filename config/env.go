package config

import "github.com/Viventio/legos/envloader"

// NOTE: "direnv" package is used for loading environment variables
// Make sure "direnv" in your machine

var CFG *EnvConfig = InitEnvConfig()

type EnvConfig struct {
	Port      int    `env:"PORT" required:"true" default:"8080"`
	MongoUri  string `env:"MONGODB_URI" required:"true"`
	JWTSecret string `env:"JWT_SECRET" required:"true"`
}

func InitEnvConfig() *EnvConfig {
	cfg := EnvConfig{}
	err := envloader.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
