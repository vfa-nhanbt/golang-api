package controllers

import (
	"os"

	"github.com/vfa-nhanbt/todo-api/db/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

var authController *AuthController

func GetAuthController() *AuthController {
	return authController
}

func InitControllers(client *mongo.Client) {
	db := os.Getenv("MONGODB_DATABASE")
	mongoDB := client.Database(db)

	/// Init auth controller
	authRepo := &repositories.UserRepository{
		MongoCollection: mongoDB.Collection("users"),
	}
	authController = &AuthController{
		Repository: authRepo,
	}
}
