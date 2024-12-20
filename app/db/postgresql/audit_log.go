package postgresql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/middleware"
	"gorm.io/gorm"
)

const (
	EventCreate = "create"
	EventUpdate = "update"
	EventDelete = "delete"
)

func logEventWithUser(tx *gorm.DB, event string) {
	ctx := tx.Statement.Context
	userID, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		fmt.Println("Cannot get userID from context")
		return
	}
	if tx.Statement.Schema != nil {
		tableName := tx.Statement.Schema.Table
		primaryKey := tx.Statement.Schema.PrioritizedPrimaryField.Name
		recordID := tx.Statement.ReflectValue.FieldByName(primaryKey).Interface()
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			fmt.Println("Cannot parse userID to uuid")
			return
		}
		recordUUID, ok := recordID.(uuid.UUID)
		if !ok {
			fmt.Println("Cannot parse recordID to uuid")
			return
		}
		log := &models.LogModel{
			UserID:        &userUUID,
			Table:         tableName,
			RecordID:      recordUUID,
			OperationType: event,
			Details:       fmt.Sprintf("Event %s by userID %s occurred on table %s for record ID %s", event, userID, tableName, recordID),
		}
		// tx.Session(&gorm.Session{}).Create(&log) // Insert log
		if err := tx.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Create(log).Error; err != nil {
			fmt.Print(fmt.Errorf("error in audit log creation: %s", err.Error()))
			return
		}
		fmt.Println("Log created successfully")
	}
}

func RegisterCallback(db *gorm.DB) {
	db.Callback().Create().After("gorm:after_create").Register("audit_log:after_create", func(db *gorm.DB) {
		logEventWithUser(db, EventCreate)
	})
	db.Callback().Update().After("gorm:after_update").Register("audit_log:after_update", func(db *gorm.DB) {
		logEventWithUser(db, EventUpdate)
	})
	db.Callback().Delete().After("gorm:after_delete").Register("audit_log:after_delete", func(db *gorm.DB) {
		logEventWithUser(db, EventDelete)
	})
}
