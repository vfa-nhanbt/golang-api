package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vfa-nhanbt/todo-api/app/models"
)

type EmailRepository struct{}

func (e *EmailRepository) SendEmail(email *models.CreateBookEmailModel) error {
	url := os.Getenv("MAILTRAP_URL")
	method := "POST"
	hostEmail := os.Getenv("HOST_EMAIL")
	templateId := os.Getenv("CREATE_BOOK_EMAIL_TEMPLATE_ID")

	payloadData := map[string]interface{}{
		"from": map[string]string{
			"email": hostEmail,
			"name":  "New book is coming~",
		},
		"to": []map[string]string{
			{
				"email": email.ReceiverEmail,
			},
		},
		"template_uuid": templateId,
		"template_variables": map[string]string{
			"authorName": email.AuthorName,
			"bookName":   email.BookTitle,
			"bookLink":   email.BookUrl,
		},
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payloadBytes))
	if err != nil {
		fmt.Println(err)
		return err
	}

	apiToken := os.Getenv("MAILTRAP_API_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
