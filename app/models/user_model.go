package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" db:"id" json:"id" validate:"uuid"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	Email        string     `gorm:"size:255;uniqueIndex" db:"email" json:"email" validate:"email,lte=255"`
	Name         string     `db:"name" json:"name"`
	PasswordHash string     `db:"password_hash" json:"password_hash,omitempty" validate:"lte=255"`
	UserRole     string     `db:"user_role" json:"user_role" validate:"userRole,lte=25"`
}

func (*UserModel) TableName() string {
	return "user_models"
}

type AdminModel struct {
	UserID uuid.UUID `gorm:"type:uuid;unique,primaryKey" db:"user_id" json:"user_id" validate:"required"`
	*UserModel
}

func (*AdminModel) TableName() string {
	return "admin_models"
}

type AuthorModel struct {
	UserID uuid.UUID `gorm:"type:uuid;unique,primaryKey" db:"user_id" json:"user_id" validate:"required"`
	*UserModel
	Books []BookModel `gorm:"foreignKey:AuthorID;references:ID"`
}

func (*AuthorModel) TableName() string {
	return "author_models"
}

type ViewerModel struct {
	UserID uuid.UUID `gorm:"type:uuid;unique,primaryKey" db:"user_id" json:"user_id" validate:"required"`
	*UserModel
	FollowedAuthors []AuthorModel `gorm:"many2many:followed_authors"`
	PaidBooks       []BookModel   `gorm:"many2many:paid_books"`
	Reviews         []ReviewModel `gorm:"foreignKey:AuthorID"`
}

func (*ViewerModel) TableName() string {
	return "viewer_models"
}
