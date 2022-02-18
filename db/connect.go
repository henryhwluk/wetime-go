package db

import (
	"context"
	"fmt"
	"log"
	"wetime-go/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root@137.220.61.135:27077/").SetAuth(
		options.Credential{
			AuthSource:    "wetime",
			AuthMechanism: "SCRAM-SHA-256",
			Username:      "henry",
			Password:      "password",
		})

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database(config.Conf.DB.Name)
	fmt.Println("Connected to MongoDB!")
}
