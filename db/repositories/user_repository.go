package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"gorm.io/gorm"
)

// type UserRepository struct {
// 	MongoCollection *mongo.Collection
// }

// func (r *UserRepository) FindUserByID(id string) (*models.UserModel, error) {
// 	userModel := models.UserModel{}

// 	// Creates a query filter to match documents in which the "id" equal to the parameter
// 	filter := bson.D{{Key: "user_id", Value: id}}
// 	err := r.MongoCollection.FindOne(context.Background(), filter).Decode(&userModel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userModel, nil
// }

// func (r *UserRepository) FindUserByEmail(email string) (*models.UserModel, error) {
// 	userModel := models.UserModel{}

// 	// Creates a query filter to match documents in which the "id" equal to the parameter
// 	filter := bson.D{{Key: "email", Value: email}}
// 	err := r.MongoCollection.FindOne(context.Background(), filter).Decode(&userModel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userModel, nil
// }

// func (r *UserRepository) InsertUser(user *models.UserModel) (interface{}, error) {
// 	res, err := r.MongoCollection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindUserByID(id string) (*models.UserModel, error) {
	userModel := models.UserModel{}
	err := r.DB.First(&userModel, id).Find(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.UserModel, error) {
	userModel := models.UserModel{}
	err := r.DB.Where("email = ?", "email").First(&userModel).Find(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (r *UserRepository) InsertUser(user *models.UserModel) error {
	err := r.DB.Create(&user).Error
	return err
}
