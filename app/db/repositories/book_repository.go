package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) InsertBook(book *models.BookModel) error {
	err := r.DB.Create(&book).Error
	return err
}

func (r *BookRepository) UpdateBook(book *models.BookModel, fieldUpdate map[string]interface{}) error {
	err := r.DB.Model(&book).Updates(fieldUpdate).Error
	return err
}

func (r *BookRepository) DeleteBookWithID(id string) error {
	err := r.DB.Delete(&models.BookModel{}, "id = ?", id).Error
	return err
}

func (r *BookRepository) GetBookByID(id string) (*models.BookModel, error) {
	bookModel := models.BookModel{}
	err := r.DB.First(&bookModel, "id = ?", id).Preload("Author").Preload("Reviews").Find(&bookModel).Error
	if err != nil {
		return nil, err
	}
	return &bookModel, nil
}

func (r *BookRepository) GetAllBooks() ([]*models.BookModel, error) {
	var books []*models.BookModel
	err := r.DB.Preload("Author").Preload("Reviews").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookByPage(page int, limit int) ([]*models.BookModel, error) {
	var books []*models.BookModel
	err := r.DB.Preload("Author").Preload("Reviews").Scopes(helpers.NewPagination(page, limit).PaginatedResult).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
