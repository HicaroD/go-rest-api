package nrdm

import (
	"context"
	"fmt"

	// MongoDB
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDatabase struct {
	client *mongo.Client
}

func ConnectToMongoDB(mongoDbUri string) (*MongoDatabase, error) {
	if mongoDbUri == "" {
		return nil, fmt.Errorf("mongodb uri should not be empty")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		return nil, err
	}

	database := &MongoDatabase{client}
	return database, nil
}

func (m *MongoDatabase) Disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
