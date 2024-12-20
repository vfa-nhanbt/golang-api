package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogModel struct {
	gorm.Model
	Table         string     `db:"table" json:"table"`
	OperationType string     `db:"operation_type" json:"operation_type"`
	RecordID      uuid.UUID  `gorm:"index"`
	UserID        *uuid.UUID `db:"user_id" json:"user_id"`
	Details       string     `db:"details" json:"details"`
}

func (*LogModel) TableName() string {
	return "audit_log_model"
}
