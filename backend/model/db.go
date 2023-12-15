package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnectToMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://todo12.mongo.cosmos.azure.com:10255/?ssl=true&retryWrites=false")
	clientOptions.SetAuth(options.Credential{
		Username: "todo12",
		Password: "SE9LyzkytHMfuwm67pwKnbLWA4iBpnIGK1Nmrd0nRRixI0ffD22Vvc7KrXofmWN8h2c8a1onJ6NLACDbzjHzBA==",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db = client.Database("todo-app") // Replace "your-database-name" with your actual database name
	log.Println("Connected to MongoDB!")
	return nil
}
