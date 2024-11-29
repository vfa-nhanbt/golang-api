package db

import (
	"context"
	"log"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	PostgresGormDB *gorm.DB
	MongoDB        *mongo.Client
}

func ConnectToDB() (*DBClient, error) {
	// return connectToMongoDB()
	return connectToPostgres()
}

func connectToMongoDB() (*DBClient, error) {
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
	return &DBClient{
		MongoDB: client,
	}, nil
}

func connectToPostgres() (*DBClient, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBClient{
		PostgresGormDB: db,
	}, nil
}
