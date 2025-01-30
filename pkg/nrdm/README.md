# NRDM - Non-Relational Database Manager

## Drivers

### MongoDB

<details>
1. Add the following imports:

```go
import (
    "fmt"
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

2. Replace the placeholder method:

```go
type MongoConn struct {
	conn *mongo.Client
}

func connectMongoDB(config DatabaseConfig) (*MongoConn, error) {
	if config.uri == "" {
		return nil, fmt.Errorf("mongodb uri must not be empty")
	}
	clientOptions := options.Client().ApplyURI(config.uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, fmt.Errorf("unable to ping MongoDB for checking connection")
	}

	conn := &MongoConn{client}
	return conn, nil
}
```

</details>
