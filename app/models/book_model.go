package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookModel struct {
	gorm.Model
	BookID       uuid.UUID  `gorm:"type:uuid;primaryKey" db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	Title        string     `db:"title" json:"name" validate:"required"`
	Description  string     `db:"description" json:"description"`
	Views        int        `db:"views" json:"views" validate:"required"`
	PasswordHash string     `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserRole     string     `db:"user_role" json:"user_role" validate:"required,userRole,lte=25"`
	AuthorID     uuid.UUID  `db:"author_id" json:"author_id" validate:"required"`
	Author       AuthorModel
}
