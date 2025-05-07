package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI isn't set in .env file")
	}
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("Error to connect to MongoDB: ", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	fmt.Println("Connected to MongoDB successfully")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("Shopper_online").Collection(collectionName)
}
