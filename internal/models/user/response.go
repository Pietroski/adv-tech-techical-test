package userModel

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	TableID   int64     `json:"tableID"`
	UserID    uuid.UUID `json:"userID"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}
