package nrdm

type Driver string

const (
	DRIVER_MONGO Driver = "mongo"
)

// DatabaseConfig holds the configuration details for non-relational databases
type DatabaseConfig struct {
	driver Driver
	uri    string
}

// Initialize MongoDB configuration
func InitMongoConfig(uri string) DatabaseConfig {
	return DatabaseConfig{
		driver: DRIVER_MONGO,
		uri:    uri,
	}
}
