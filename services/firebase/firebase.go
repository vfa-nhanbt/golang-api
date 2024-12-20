package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var firebaseClient *firebase.App

func InitFirebaseClient() error {
	opt := option.WithCredentialsFile("firebase-golang-api-service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	firebaseClient = app
	return nil
}

func GetFirebaseMessages() (*messaging.Client, error) {
	firebaseMsg, err := firebaseClient.Messaging(context.Background())
	if err != nil {
		return nil, err
	}
	return firebaseMsg, nil
}
