package models

import (
	"time"

	"github.com/google/uuid"
)

type BookModel struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey" db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   *time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time    `db:"updated_at" json:"updated_at"`
	Title       string        `db:"title" json:"title" validate:"required"`
	Description string        `db:"description" json:"description"`
	Views       int           `gorm:"default:0" db:"views" json:"views"`
	Price       int           `gorm:"default:0" db:"price" json:"price"`
	AuthorID    uuid.UUID     `db:"author_id" json:"author_id" validate:"required"`
	Author      AuthorModel   `gorm:"foreignKey:AuthorID;references:ID"`
	Reviews     []ReviewModel `gorm:"foreignKey:BookID"`
}

func (*BookModel) TableName() string {
	return "book_models"
}
