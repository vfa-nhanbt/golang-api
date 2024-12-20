package models

import "firebase.google.com/go/v4/messaging"

type NotificationModel struct {
	Title        string            `json:"title" validate:"required"`
	Body         string            `json:"body" validate:"required"`
	Data         map[string]string `json:"data,omitempty"`
	Topic        string            `json:"topic,omitempty"`
	DeviceTokens []string          `json:"device_tokens,omitempty"`
	Condition    string            `json:"condition,omitempty"`
}

func (n *NotificationModel) ToFirebaseMessage() *messaging.Message {
	return &messaging.Message{
		Data:  n.Data,
		Topic: n.Topic,
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Token:     n.DeviceTokens[0],
		Condition: n.Condition,
	}
}

func (n *NotificationModel) ToFirebaseMessageWithMultipleTokens() *messaging.MulticastMessage {
	return &messaging.MulticastMessage{
		Data: n.Data,
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Tokens: n.DeviceTokens,
	}
}

type SubscribeToTopicModel struct {
	RegistrationTokens []string `json:"registration_tokens" validate:"required"`
	Topic              string   `json:"topic" validate:"required"`
}
