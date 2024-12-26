package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/vfa-nhanbt/todo-api/app/models"
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
	/// Get the database connection settings from env
	// host := os.Getenv("POSTGRES_HOST")
	// username := os.Getenv("POSTGRES_USERNAME")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// port := os.Getenv("POSTGRES_PORT")
	// dbname := os.Getenv("POSTGRES_NAME")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, dbname, port)
	databaseUrl := os.Getenv("POSTGRES_RAILWAY_URL")
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	fmt.Println("Connected to the database successfully!")

	return &DBClient{
		PostgresGormDB: db,
	}, nil
}

func PostgresAutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		/// Migrate user models
		&models.UserModel{},
		/// Migrate admin models
		&models.AdminModel{},
		/// Migrate author models
		&models.AuthorModel{},
		/// Migrate viewer models
		&models.ViewerModel{},
		/// Migrate books models
		&models.BookModel{},
		/// Migrate review models
		&models.ReviewModel{},
		/// Migrate audit log models
		&models.LogModel{},
	)
	if err != nil {
		return fmt.Errorf("cannot migrate table with error: %v", err)
	}
	return nil
}
