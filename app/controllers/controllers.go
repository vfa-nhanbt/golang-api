package controllers

import (
	"github.com/vfa-nhanbt/todo-api/db"
	"github.com/vfa-nhanbt/todo-api/db/repositories"
)

var authController *AuthController

func GetAuthController() *AuthController {
	return authController
}

func InitControllers(db *db.DBClient) {
	// mongo := os.Getenv("MONGODB_DATABASE")
	// mongoDB := db.MongoDB.Database(mongo)

	/// Init auth controller
	authRepo := &repositories.UserRepository{
		DB: db.PostgresGormDB,
	}

	authController = &AuthController{
		Repository: authRepo,
	}
}
