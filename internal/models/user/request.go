package userModel

import "github.com/google/uuid"

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type UpdateUserRequest struct {
	UserID    uuid.UUID `json:"userID"`
	Email     string    `json:"email" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
}

type UserQueryParams struct {
	UserID string `form:"user_id" binding:"omitempty,uuid"`
	Email  string `form:"email" binding:"omitempty,email"`
}

type PaginationQueryParams struct {
	PageID   int32 `form:"page_id" binding:"omitempty,numeric,min=1"`
	PageSize int32 `form:"page_size" binding:"omitempty,numeric,min=1"`
	// Limit    string `form:"limit"`  // req.PageSize
	// Offset   string `form:"offset"` // (req.PageID - 1) * req.PageSize
}
