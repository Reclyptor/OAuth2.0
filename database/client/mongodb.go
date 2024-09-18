package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewMongoDBClient() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("MONGODB_URI is required!")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}
