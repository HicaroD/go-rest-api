package rdm

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
