package connect

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection URI from environment variables
var mongoURI string = os.Getenv("MONGODB_URI")

// ConnectDB initializes a MongoDB client and returns a connection
func ConnectDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	log.Println("Connected to MongoDB")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	database := os.Getenv("MONGODB_DATABASE") // Database name from environment variables
	collection := client.Database(database).Collection(collectionName)
	return collection
}
