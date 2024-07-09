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
	clientOptions := options.Client().ApplyURI("mongodb://website.mongo.cosmos.azure.com:10255/?ssl=true&retryWrites=false")
	clientOptions.SetAuth(options.Credential{
		Username: "website",
		Password: "8biA4kRmOESvlIDN4IM3gqnqFK69xOZHEsQJtgqCRgzhDzFbEj9EdRYVr0pbOd2EEptdPACfjiwHACDbSsHCGQ==",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
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
