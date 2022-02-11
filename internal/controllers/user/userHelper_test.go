package userController_test

import (
	"database/sql"
	userModel "github.com/Pietroski/adv-tech-techical-test/internal/models/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func randomCreateUserRequest() *userModel.CreateUserRequest {
	ur := &userModel.CreateUserRequest{
		Email:     "some-email@valid.co",
		FirstName: "some",
		LastName:  "email",
	}

	return ur
}

func createUserParams(u *userModel.CreateUserRequest) *user.CreateUserParams {
	uResp := &user.CreateUserParams{
		UserID:    uuid.UUID{},
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	return uResp
}

func createUserResult(
	u *user.CreateUserParams,
	tableID int64,
	t time.Time,
) *user.Users {
	uResp := &user.Users{
		TableID: sql.NullInt64{
			Int64: tableID,
			Valid: true,
		},
		UserID:    u.UserID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: sql.NullTime{
			Time:  t,
			Valid: true,
		},
	}

	return uResp
}

func ExpectedUserResult(
	u *user.Users,
	tableID int64,
	t time.Time,
) *userModel.UserResponse {
	eur := &userModel.UserResponse{
		TableID:   tableID,
		UserID:    u.UserID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: t,
	}

	return eur
}

func createRandomUserList(
	t *testing.T,
	size int,
) (
	[]user.Users,
	[]userModel.UserResponse,
) {
	xcup := []user.Users{}
	xur := []userModel.UserResponse{}

	for i := 0; i < size; i++ {
		userUUID, err := uuid.NewUUID()
		require.NoError(t, err)
		tn := time.Now()
		ti := int64(4)

		cur := createUserResult(&user.CreateUserParams{
			UserID:    userUUID,
			Email:     "some-email",
			FirstName: "some",
			LastName:  "email",
		}, ti, tn)
		xcup = append(xcup, *cur)
		eur := ExpectedUserResult(cur, ti, tn)
		xur = append(xur, *eur)
	}

	return xcup, xur
}
