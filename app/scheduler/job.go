package scheduler

import (
	"fmt"

	"github.com/vfa-nhanbt/todo-api/app/db"
	"github.com/vfa-nhanbt/todo-api/app/db/repositories"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/services/firebase"
)

type SendNotificationJob struct {
	Service  *firebase.FirebaseMessagingService
	BookRepo *repositories.BookRepository
	UserRepo *repositories.UserRepository
}

func InitSendNotificationJob(db *db.DBClient) *SendNotificationJob {
	return &SendNotificationJob{
		Service:  &firebase.FirebaseMessagingService{},
		BookRepo: &repositories.BookRepository{DB: db.PostgresGormDB},
		UserRepo: &repositories.UserRepository{DB: db.PostgresGormDB},
	}
}

func (job *SendNotificationJob) SendBookNotificationToUser() error {
	/// Find user who has registered device token
	users, err := job.UserRepo.FindUserHasDeviceToken()
	if err != nil {
		return err
	}
	var tokens []string
	for _, user := range users {
		tokens = append(tokens, user.DeviceTokens...)
	}
	randomBook, err := job.BookRepo.GetRandomBook()
	if err != nil {
		return err
	}

	notification := &models.NotificationModel{
		Title:        "Book for today!",
		Body:         fmt.Sprintf("Hello, Start new day with %s from %s", randomBook.Title, randomBook.Author.Name),
		DeviceTokens: tokens,
	}

	err = job.Service.SendNotificationWithMultipleTokens(notification)
	return err
}
