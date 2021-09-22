package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

// connect to mongodb
func ConnectDB() error {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return err
	}

	MI = MongoInstance{
		Client: client,
		DB:     client.Database("todos"),
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}
