package nrdm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectMongoDB initializes the connection for MongoDB.
// func connectMongoDB(config DatabaseConfig) (*mongo.Client, error) {
// 	return nil, fmt.Errorf("mongodb driver not supported")
// }

// connectMongoDB initializes the connection for MongoDB.
func connectMongoDB(config DatabaseConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return client, nil
}
