// Package nrdm (Non-Relational Database Manager) provides some convenient
// methods for connecting to non-relational databases.
package nrdm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct encapsulates the MongoDB connection and its methods
type Database struct {
	config DatabaseConfig
	Conn   *mongo.Client
}

// Connect initializes the connection to the non-relational database based on the driver
func Connect(config DatabaseConfig) (*Database, error) {
	var err error
	db := &Database{}

	switch config.driver {
	case DRIVER_MONGO:
		db.Conn, err = connectMongoDB(config)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.config = config
	return db, nil
}

func (db *Database) Disconnect() {
	switch db.config.driver {
	case DRIVER_MONGO:
		if err := db.Conn.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}
