package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) InsertBook(book *models.BookModel) error {
	err := r.DB.Create(&book).Error
	return err
}

func (r *BookRepository) UpdateBook(book *models.BookModel) error {
	// err := r.DB.Save(&book).Error	/// Update all fields
	err := r.DB.Model(&book).Updates(book).Error
	return err
}

func (r *BookRepository) DeleteBookWithID(id string) error {
	err := r.DB.Delete(&models.BookModel{}, id).Error
	return err
}

func (r *BookRepository) GetBookByID(id string) (*models.BookModel, error) {
	bookModel := models.BookModel{}
	err := r.DB.First(&bookModel, id).Preload("Author").Preload("Reviews").Find(&bookModel).Error
	if err != nil {
		return nil, err
	}
	return &bookModel, nil
}

func (r *BookRepository) GetAllBooks() ([]*models.BookModel, error) {
	var books []*models.BookModel
	err := r.DB.Find(&books).Preload("Author").Preload("Reviews").Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) CheckIsAuthor(authorId string, bookId string) (bool, error) {
	return true, nil
}
