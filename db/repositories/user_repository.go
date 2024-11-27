package repositories

import (
	"context"

	"github.com/vfa-nhanbt/todo-api/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	MongoCollection *mongo.Collection
}

func (r *UserRepository) FindUserByID(id string) (*models.UserModel, error) {
	userModel := models.UserModel{}

	// Creates a query filter to match documents in which the "id" equal to the parameter
	filter := bson.D{{Key: "user_id", Value: id}}
	err := r.MongoCollection.FindOne(context.Background(), filter).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.UserModel, error) {
	userModel := models.UserModel{}

	// Creates a query filter to match documents in which the "id" equal to the parameter
	filter := bson.D{{Key: "email", Value: email}}
	err := r.MongoCollection.FindOne(context.Background(), filter).Decode(&userModel)
	if err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepository) InsertUser(user *models.UserModel) (interface{}, error) {
	res, err := r.MongoCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
