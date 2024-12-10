package models

import (
	"time"

	"github.com/google/uuid"
)

type ReviewModel struct {
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey" db:"id" json:"id" validate:"required,uuid"`
	Description string      `db:"description" json:"description" validate:"required"`
	ViewerID    uuid.UUID   `db:"viewer_id" json:"viewer_id" validate:"required"`
	Viewer      ViewerModel `gorm:"foreignKey:ViewerID;references:ID"`
	BookID      uuid.UUID   `db:"book_id" json:"book_id" validate:"required"`
	Book        BookModel   `gorm:"foreignKey:BookID"`
	CreatedAt   *time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time  `db:"updated_at" json:"updated_at"`
}

func (*ReviewModel) TableName() string {
	return "review_models"
}
