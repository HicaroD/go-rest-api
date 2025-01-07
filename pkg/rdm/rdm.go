// Package rdm (Relational Database Manager) provides some convenient methods
// for connecting to databases using GORM.
package rdm

import (
	"fmt"
	"gorm.io/gorm"
	// Uncomment this for SQLite
	"gorm.io/driver/sqlite"
	// Uncomment this for MySQL
	// "gorm.io/driver/mysql"
	// Uncomment this PostgreSQL
	// "gorm.io/driver/postgres"
)

type Driver string

const (
	DRIVER_MYSQL    Driver = "mysql"
	DRIVER_POSTGRES Driver = "postgres"
	DRIVER_SQLITE   Driver = "sqlite"
)

// DatabaseConfig holds the database configuration details
type DatabaseConfig struct {
	Driver   Driver
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Initialize database configuration for SQLite only
// Consider using InitDatabaseConfig method for other database systems.
func InitSqliteConfig(path string) DatabaseConfig {
	return DatabaseConfig{Driver: DRIVER_SQLITE, DBName: path}
}

// Initialize database configuration for MySQL and PostgreSQL
func InitDatabaseConfig(driver Driver, host, port, user, password, dbName string) DatabaseConfig {
	return DatabaseConfig{
		Driver:   driver,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

// Database struct encapsulates the connection and its methods
type Database struct {
	Conn *gorm.DB
}

// Connect initializes the database connection
func Connect(config DatabaseConfig) (*Database, error) {
	db := &Database{}

	dsn, err := buildDSN(config)
	if err != nil {
		return nil, err
	}

	switch config.Driver {
	case DRIVER_SQLITE:
		db.Conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	// case MYSQL_DRIVER:
	// 	db.Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// case POSTGRES_DRIVER:
	// 	db.Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// buildDSN constructs the DSN string based on the configuration
func buildDSN(config DatabaseConfig) (string, error) {
	switch config.Driver {
	case DRIVER_SQLITE:
		return config.DBName, nil
	case DRIVER_MYSQL:
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User, config.Password, config.Host, config.Port, config.DBName), nil
	case DRIVER_POSTGRES:
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DBName), nil
	default:
		return "", fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}

func (db *Database) AutoMigrate(models ...interface{}) error {
	if db.Conn == nil {
		return fmt.Errorf("database is not connected")
	}
	return db.Conn.AutoMigrate(models...)
}
