package repositories

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) InsertBook(book *models.BookModel, ctx ...context.Context) error {
	err := helpers.GetDB(r.DB, ctx...).Create(&book).Error
	return err
}

func (r *BookRepository) UpdateBook(book *models.BookModel, fieldUpdate map[string]interface{}, ctx ...context.Context) error {
	err := helpers.GetDB(r.DB, ctx...).Model(&book).Updates(fieldUpdate).Error
	return err
}

func (r *BookRepository) DeleteBookWithID(id string, ctx ...context.Context) error {
	err := helpers.GetDB(r.DB, ctx...).Delete(&models.BookModel{}, "id = ?", id).Error
	return err
}

func (r *BookRepository) GetBookByID(id string, ctx ...context.Context) (*models.BookModel, error) {
	bookModel := models.BookModel{}
	err := helpers.GetDB(r.DB, ctx...).First(&bookModel, "id = ?", id).Preload("Author").Preload("Reviews").Find(&bookModel).Error
	if err != nil {
		return nil, err
	}
	return &bookModel, nil
}

func (r *BookRepository) GetAllBooks(ctx ...context.Context) ([]*models.BookModel, error) {
	var books []*models.BookModel
	err := helpers.GetDB(r.DB, ctx...).Preload("Author").Preload("Reviews").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookByPage(page int, limit int, ctx ...context.Context) ([]*models.BookModel, error) {
	var books []*models.BookModel
	err := helpers.GetDB(r.DB, ctx...).Preload("Author").Preload("Reviews").Scopes(helpers.NewPagination(page, limit).PaginatedResult).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetRandomBook(ctx ...context.Context) (*models.BookModel, error) {
	var count int64
	err := helpers.GetDB(r.DB, ctx...).Model(&models.BookModel{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	rand.NewSource(time.Now().UnixNano())
	randomOffset := rand.Int63n(count)
	book := models.BookModel{}
	err = helpers.GetDB(r.DB, ctx...).Preload("Author").Preload("Reviews").Offset(int(randomOffset)).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) SearchBooks(page int, limit int, query string, ctx ...context.Context) ([]*models.BookModel, error) {
	var books []*models.BookModel
	words := strings.Fields(query)
	queryCommand := strings.Join(words, " & ")
	err := helpers.GetDB(r.DB, ctx...).
		Preload("Author").Preload("Reviews").
		Scopes(helpers.NewPagination(page, limit).PaginatedResult).
		Where("search_vector @@ to_tsquery('english', ?)", queryCommand).
		Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}
