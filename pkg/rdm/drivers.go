package rdm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// connectSQLite initializes the connection for SQLite.
func connectSQLite(config DatabaseConfig) (*gorm.DB, error) {
	dsn := config.DBName
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

// connectPostgres initializes the connection for PostgreSQL.
func connectPostgres(config DatabaseConfig) (*gorm.DB, error) {
	return nil, fmt.Errorf("postgresql driver is not supported.")
}

// connectMySQL initializes the connection for MySQL.
func connectMySQL(config DatabaseConfig) (*gorm.DB, error) {
	return nil, fmt.Errorf("mysql driver is not supported.")
}
