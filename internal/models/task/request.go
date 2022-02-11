package taskModel

import (
	"database/sql"
	"github.com/google/uuid"
)

type TaskRequest struct {
	UserID         uuid.UUID     `json:"userID" binding:"required"`
	DataRange      interface{}   `json:"dataRange" binding:"required"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod" binding:"required"`
}
