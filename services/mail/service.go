package mail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type EmailService struct{}

func (e *EmailService) SendEmail(email EmailTemplate) error {
	url := os.Getenv("MAILTRAP_URL")
	method := "POST"

	payloadBytes, err := json.Marshal(email.GeneratePayload())
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

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var responseMap map[string]interface{}
		if err := json.Unmarshal(body, &responseMap); err == nil {
			if error, ok := responseMap["errors"].([]interface{}); ok && len(error) > 0 {
				if errMsg, ok := error[0].(string); ok {
					return errors.New(errMsg)
				}
			}
		}
		return errors.New(string(body))
	}
	fmt.Println(string(body))
	return nil
}
