package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection

func ConnectDatabase() {
	var err error
	mongoURI := os.Getenv("MONGODB_URI")
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = Client.Database("gin-boilerplate").Collection("users")
}
