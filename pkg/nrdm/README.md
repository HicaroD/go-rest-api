# NRDM - Non-Relational Database Manager

## Drivers

### MongoDB

<details>
1. Add the following imports:

```go
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

2. Replace the placeholder method:

```go
// connectMongoDB initializes the connection for MongoDB.
func connectMongoDB(config DatabaseConfig) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(config.uri)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    return client, nil
}
```

</details>
