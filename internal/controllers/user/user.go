package userController

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	userModel "github.com/Pietroski/adv-tech-techical-test/internal/models/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/util/notification"
)

type (
	Server struct {
		store user.Store
	}
)

func NewServer(store user.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var (
		req userModel.CreateUserRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	newUserID := uuid.New()
	createUserPayload := user.CreateUserParams{
		UserID:    newUserID,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	createdUser, err := s.store.CreateUser(ctx, createUserPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	cur := populateUserResponse(createdUser)
	ctx.JSON(http.StatusCreated, cur)
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	var (
		req userModel.UpdateUserRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	updateUserPayload := user.UpdateUserParams{
		UserID:    req.UserID,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	updatedUser, err := s.store.UpdateUser(ctx, updateUserPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	uur := populateUserResponse(updatedUser)
	ctx.JSON(http.StatusOK, uur)
}

func (s *Server) GetUser(ctx *gin.Context) {
	var (
		req userModel.UserQueryParams
	)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	if req.UserID != "" && req.Email != "" {
		log.Println(req.Email)
		errMsg := errors.New("cannot use both user_id and email for querying")
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(errMsg))
		return
	} else if req.UserID == "" && req.Email == "" {
		errMsg := errors.New("invalid query parameter for a single user")
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(errMsg))
		return
	}

	if req.UserID != "" {
		s.getUserById(ctx, req)
		return
	}

	s.getUserByEmail(ctx, req)
	return
}

func (s *Server) getUserById(
	ctx *gin.Context,
	req userModel.UserQueryParams,
) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	u, err := s.store.GetUserByID(ctx, userUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	uur := populateUserResponse(u)
	ctx.JSON(http.StatusOK, uur)
}

func (s *Server) getUserByEmail(
	ctx *gin.Context,
	req userModel.UserQueryParams,
) {
	u, err := s.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	uur := populateUserResponse(u)
	ctx.JSON(http.StatusOK, uur)
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	var (
		req userModel.UserQueryParams
	)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	if req.UserID != "" && req.Email != "" {
		log.Println(req.Email)
		errMsg := errors.New("cannot use both user_id and email for querying")
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(errMsg))
		return
	} else if req.UserID == "" && req.Email == "" {
		errMsg := errors.New("invalid query parameter for a single user")
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(errMsg))
		return
	}

	if req.UserID != "" {
		s.deleteUserByID(ctx, req)
		return
	}

	s.deleteUserByEmail(ctx, req)
	return
}

func (s *Server) deleteUserByID(
	ctx *gin.Context,
	req userModel.UserQueryParams,
) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	err = s.store.DeleteUserByID(ctx, userUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(
			"user with uuid: %v - deleted successfully",
			userUUID.String(),
		),
	})
}

func (s *Server) deleteUserByEmail(
	ctx *gin.Context,
	req userModel.UserQueryParams,
) {
	err := s.store.DeleteUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(
			"user with email: %v - deleted successfully",
			req.Email,
		),
	})
}

func (s *Server) ListUsers(ctx *gin.Context) {
	var (
		req userModel.PaginationQueryParams
	)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	if req.PageID != 0 && req.PageSize != 0 {
		s.listPaginatedUsers(ctx, req)
		return
	} else if req.PageID == 0 && req.PageSize == 0 {
		s.listUsers(ctx)
		return
	}

	errMsg := errors.New("cannot provide only one of the pagination elements")
	ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(errMsg))
	return
}

func (s *Server) listUsers(
	ctx *gin.Context,
) {
	xu, err := s.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	pul := populateUserResponseList(xu)
	ctx.JSON(http.StatusOK, pul)
}

func (s *Server) listPaginatedUsers(
	ctx *gin.Context,
	req userModel.PaginationQueryParams,
) {
	pagParams := user.ListPaginatedUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	xu, err := s.store.ListPaginatedUsers(ctx, pagParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	pul := populateUserResponseList(xu)
	ctx.JSON(http.StatusOK, pul)
}

func populateUserResponse(u user.Users) userModel.UserResponse {
	return userModel.UserResponse{
		TableID:   u.TableID.Int64,
		UserID:    u.UserID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt.Time,
	}
}

func populateUserResponseList(xu []user.Users) []userModel.UserResponse {
	pul := []userModel.UserResponse{}
	for _, u := range xu {
		pul = append(pul, populateUserResponse(u))
	}
	return pul
}
