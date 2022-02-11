package taskModel

import (
	"database/sql"
	"github.com/google/uuid"
)

type TaskResponse struct {
	TableID        sql.NullInt64 `json:"tableID"`
	TaskID         uuid.UUID     `json:"taskID"`
	UserID         uuid.UUID     `json:"userID"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
	CreatedAt      sql.NullTime  `json:"createdAt"`
}
