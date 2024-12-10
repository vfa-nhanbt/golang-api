package controllers

import (
	"github.com/vfa-nhanbt/todo-api/db"
	"github.com/vfa-nhanbt/todo-api/db/repositories"
)

var authController *AuthController
var bookController *BookController
var reviewController *ReviewController
var sendEmailController *SendEmailController

func GetAuthController() *AuthController {
	return authController
}
func GetBookController() *BookController {
	return bookController
}
func GetReviewController() *ReviewController {
	return reviewController
}
func GetSendEmailController() *SendEmailController {
	return sendEmailController
}

func InitControllers(db *db.DBClient) {
	// mongo := os.Getenv("MONGODB_DATABASE")
	// mongoDB := db.MongoDB.Database(mongo)

	/// Init auth repository
	authRepo := &repositories.UserRepository{
		DB: db.PostgresGormDB,
	}
	/// Init book repository
	bookRepo := &repositories.BookRepository{
		DB: db.PostgresGormDB,
	}
	/// Init review repository
	reviewRepo := &repositories.ReviewRepository{
		DB: db.PostgresGormDB,
	}

	authController = &AuthController{
		Repository: authRepo,
	}
	bookController = &BookController{
		Repository: bookRepo,
	}
	reviewController = &ReviewController{
		Repository: reviewRepo,
	}
	sendEmailController = &SendEmailController{
		Repository: &repositories.EmailRepository{},
	}
}
