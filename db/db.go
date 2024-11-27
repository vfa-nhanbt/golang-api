package db

import (
	"context"
	"log"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {
	/// Get the database connection string
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	cluster := os.Getenv("MONGODB_CLUSTER")
	appname := os.Getenv("MONGODB_APPNAME")
	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster + "/?retryWrites=true&w=majority&appName=" + appname

	clientOptions := options.Client().ApplyURI(uri)

	/// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Connected to mongodb successfully..")
	return client, nil
}
