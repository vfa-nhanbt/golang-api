package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
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

// / ----- Admin -----
func (r *UserRepository) FindUserByID(id string) (*models.AdminModel, error) {
	adminModel := models.AdminModel{}
	err := r.DB.First(&adminModel, "id = ?", id).Find(&adminModel).Error
	if err != nil {
		return nil, err
	}
	return &adminModel, nil
}

func (r *UserRepository) FindAdminByEmail(email string) (*models.AdminModel, error) {
	adminModel := &models.AdminModel{}
	err := r.DB.Where("email = ?", email).First(adminModel).Find(adminModel).Error
	if err == nil {
		return adminModel, nil
	}
	if err.Error() == "record not found" {
		return nil, nil
	}
	return nil, err
}

func (r *UserRepository) InsertAdmin(user *models.AdminModel) error {
	/// Encode user password before insert to db
	decodePassword, err := helpers.EncodeUserPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = decodePassword
	err = r.DB.Create(&user).Error
	return err
}

/// ----------

// / ----- Author -----
func (r *UserRepository) FindAuthorByID(id string) (*models.AuthorModel, error) {
	authorModel := models.AuthorModel{}
	err := r.DB.First(&authorModel, "id = ?", id).Preload("Books").Find(&authorModel).Error
	if err != nil {
		return nil, err
	}
	return &authorModel, nil
}

func (r *UserRepository) FindAuthorByEmail(email string) (*models.AuthorModel, error) {
	authorModel := &models.AuthorModel{}
	err := r.DB.Where("email = ?", email).Preload("Books").First(authorModel).Find(authorModel).Error
	if err == nil {
		return authorModel, nil
	}
	if err.Error() == "record not found" {
		return nil, nil
	}
	return nil, err
}

func (r *UserRepository) InsertAuthor(user *models.AuthorModel) error {
	/// Encode user password before insert to db
	decodePassword, err := helpers.EncodeUserPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = decodePassword
	err = r.DB.Create(&user).Error
	return err
}

// func (r *UserRepository) LoadBooks(authorModel *models.AuthorModel) error {
// 	err := r.DB.Preload("Books").Where("id = ?", authorModel.AuthorID).First(&authorModel).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

/// ----------

// / ----- Viewer -----
func (r *UserRepository) FindViewerByID(id string) (*models.ViewerModel, error) {
	viewerModel := models.ViewerModel{}
	err := r.DB.First(&viewerModel, "id = ?", id).Preload("FollowedAuthors").Find(&viewerModel).Error
	if err != nil {
		return nil, err
	}
	return &viewerModel, nil
}

func (r *UserRepository) FindViewerByEmail(email string) (*models.ViewerModel, error) {
	viewerModel := &models.ViewerModel{}
	err := r.DB.Where("email = ?", email).Preload("FollowedAuthors").First(viewerModel).Find(viewerModel).Error
	if err == nil {
		return viewerModel, nil
	}
	if err.Error() == "record not found" {
		return nil, nil
	}
	return nil, err
}

func (r *UserRepository) InsertViewer(user *models.ViewerModel) error {
	/// Encode user password before insert to db
	decodePassword, err := helpers.EncodeUserPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = decodePassword
	err = r.DB.Create(&user).Error
	return err
}

// func (r *UserRepository) LoadFavoriteAuthors(viewerModel *models.ViewerModel) error {
// 	err := r.DB.Preload("FollowedAuthors").Where("id = ?", viewerModel.ViewerID).First(&viewerModel).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

/// ----------
