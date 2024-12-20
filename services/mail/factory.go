package mail

import (
	"os"

	"github.com/vfa-nhanbt/todo-api/app/models"
)

type EmailTemplate interface {
	GeneratePayload() map[string]interface{}
}

type CreateBookEmail struct {
	CreateBookModel models.CreateBookEmailModel
}

func (e CreateBookEmail) GeneratePayload() map[string]interface{} {
	templateId := os.Getenv("CREATE_BOOK_EMAIL_TEMPLATE_ID")
	hostEmail := os.Getenv("HOST_EMAIL")

	return map[string]interface{}{
		"from": map[string]string{
			"email": hostEmail,
			"name":  "New book is coming~",
		},
		"to": []map[string]string{
			{
				"email": e.CreateBookModel.ReceiverEmail,
			},
		},
		"template_uuid": templateId,
		"template_variables": map[string]string{
			"authorName": e.CreateBookModel.AuthorName,
			"bookName":   e.CreateBookModel.BookTitle,
			"bookLink":   e.CreateBookModel.BookUrl,
		},
	}
}

type WelcomeEmail struct {
	WelcomeModel *models.WelcomeEmailModel
}

func (e *WelcomeEmail) GeneratePayload() map[string]interface{} {
	templateId := os.Getenv("WELCOME_EMAIL_TEMPLATE_ID")
	hostEmail := os.Getenv("HOST_EMAIL")

	return map[string]interface{}{
		"from": map[string]string{
			"email": hostEmail,
			"name":  "Welcome to our application~",
		},
		"to": []map[string]string{
			{
				"email": e.WelcomeModel.ReceiverEmail,
			},
		},
		"template_uuid": templateId,
		"template_variables": map[string]string{
			"userName": e.WelcomeModel.UserName,
		},
	}
}
