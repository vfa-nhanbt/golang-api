package firebase

import (
	"context"
	"fmt"
	"log"

	"github.com/vfa-nhanbt/todo-api/app/models"
)

type FirebaseMessagingService struct{}

func (m *FirebaseMessagingService) SendNotification(notification *models.NotificationModel) error {
	client, err := GetFirebaseMessages()
	if err != nil {
		return err
	}
	response, err := client.Send(context.Background(), notification.ToFirebaseMessage())
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	return nil
}

func (m *FirebaseMessagingService) SendNotificationWithMultipleTokens(notification *models.NotificationModel) error {
	client, err := GetFirebaseMessages()
	if err != nil {
		return err
	}
	response, err := client.SendEachForMulticast(context.Background(), notification.ToFirebaseMessageWithMultipleTokens())
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	return nil
}

func (m *FirebaseMessagingService) SubscribeToTopic(subscribeToTopicModel *models.SubscribeToTopicModel) error {
	client, err := GetFirebaseMessages()
	if err != nil {
		return err
	}
	response, err := client.SubscribeToTopic(context.Background(), subscribeToTopicModel.RegistrationTokens, subscribeToTopicModel.Topic)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
	return nil
}

func (m *FirebaseMessagingService) UnsubscribeToTopic(subscribeToTopicModel *models.SubscribeToTopicModel) error {
	client, err := GetFirebaseMessages()
	if err != nil {
		return err
	}
	response, err := client.UnsubscribeFromTopic(context.Background(), subscribeToTopicModel.RegistrationTokens, subscribeToTopicModel.Topic)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.SuccessCount, "tokens were unsubscribed successfully")
	return nil
}
