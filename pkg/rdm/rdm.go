// Package rdm (Relational Database Manager) provides some convenient methods
// for connecting to databases using GORM.
package rdm

import (
	"fmt"
	"gorm.io/gorm"
)

// Database struct encapsulates the connection and its methods
type Database struct {
	Conn *gorm.DB
}

// Initializes the database connection based on the driver
func Connect(config DatabaseConfig) (*Database, error) {
	var err error
	db := &Database{}

	switch config.Driver {
	case DRIVER_SQLITE:
		db.Conn, err = connectSQLite(config)
	case DRIVER_MYSQL:
		db.Conn, err = connectMySQL(config)
	case DRIVER_POSTGRES:
		db.Conn, err = connectPostgres(config)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// Automigrate any number of models
func (db *Database) AutoMigrate(models ...interface{}) error {
	if db.Conn == nil {
		return fmt.Errorf("database is not connected")
	}
	return db.Conn.AutoMigrate(models...)
}

// // Constructs the DSN string based on the configuration
// func buildDSN(config DatabaseConfig) (string, error) {
// 	switch config.Driver {
// 	case DRIVER_SQLITE:
// 		return config.DBName, nil
// 	case DRIVER_MYSQL:
// 		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 			config.User, config.Password, config.Host, config.Port, config.DBName), nil
// 	case DRIVER_POSTGRES:
// 		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 			config.Host, config.Port, config.User, config.Password, config.DBName), nil
// 	default:
// 		return "", fmt.Errorf("unsupported driver: %s", config.Driver)
// 	}
// }
