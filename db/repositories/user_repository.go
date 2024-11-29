package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

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
	userModel := &models.UserModel{}
	err := r.DB.Where("email = ?", email).First(userModel).Find(userModel).Error
	if err == nil {
		return userModel, nil
	}
	if err.Error() == "record not found" {
		return nil, nil
	}
	return nil, err
}

func (r *UserRepository) InsertUser(user *models.UserModel) error {
	/// Encode user password before insert to db
	decodePassword, err := helpers.EncodeUserPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = decodePassword
	err = r.DB.Create(&user).Error
	return err
}
