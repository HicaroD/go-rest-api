package utils

import "fmt"

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func BuildConnectionString(cfg DBConfig) (string, error) {
	switch cfg.Driver {
	case "sqlite", "sqlite3":
		return fmt.Sprintf("%s.db", cfg.DBName), nil
	case "postgres":
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
		), nil
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
		), nil
	default:
		return "", fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
}
