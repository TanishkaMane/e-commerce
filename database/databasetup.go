package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use mongo.Connect to create and connect the client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://svcmstuser:svcmstuser%232022@10.26.25.154:27017/JioServiceDesk?authSource=admin&readPreference=primary"))
	if err != nil {
		log.Fatal("Error creating MongoDB client: ", err)
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB : ", err)
		return nil
	}

	fmt.Println("Successfully Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBSet()

func GetMongoCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}
