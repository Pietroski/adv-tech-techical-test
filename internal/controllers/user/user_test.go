package userController_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/Pietroski/adv-tech-techical-test/internal/factories/user"
	userModel "github.com/Pietroski/adv-tech-techical-test/internal/models/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
	mockdbcompany "github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user/mock"
)

func TestServer_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdbcompany.NewMockStore(ctrl)
	server := userFactory.NewUserFactory(store)

	uReq := randomCreateUserRequest()
	params := createUserParams(uReq)
	tn := time.Now()
	ti := int64(4)
	cur := createUserResult(params, ti, tn)
	eur := ExpectedUserResult(cur, ti, tn)
	{
		store.
			EXPECT().
			CreateUser(gomock.Any(), EqCreateCompanyParams(*params)).
			Times(1).
			Return(*cur, nil)
	}

	postReq, err := json.Marshal(uReq)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/users")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postReq))
	require.NoError(t, err)
	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusCreated, recorder.Code)
	checkResponseBody(t, eur, recorder.Body.Bytes())
}

func TestServer_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdbcompany.NewMockStore(ctrl)
	server := userFactory.NewUserFactory(store)

	userUUID, err := uuid.NewUUID()
	require.NoError(t, err)
	uur := userModel.UpdateUserRequest{
		UserID:    userUUID,
		Email:     "some-email",
		FirstName: "some",
		LastName:  "email",
	}
	params := user.UpdateUserParams(uur)
	tn := time.Now()
	ti := int64(4)
	cur := createUserResult(&user.CreateUserParams{
		UserID:    uur.UserID,
		Email:     uur.Email,
		FirstName: uur.FirstName,
		LastName:  uur.LastName,
	}, ti, tn)
	eur := ExpectedUserResult(cur, ti, tn)
	{
		store.
			EXPECT().
			UpdateUser(gomock.Any(), gomock.Eq(params)).
			Times(1).
			Return(*cur, nil)
	}

	postReq, err := json.Marshal(uur)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/users")
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(postReq))
	require.NoError(t, err)
	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	checkResponseBody(t, eur, recorder.Body.Bytes())
}

func TestServer_GetUser(t *testing.T) {
	t.Run("fails to get user - both user_id and email provided", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		userUUID, err := uuid.NewUUID()
		require.NoError(t, err)
		email := "some/email@test.com"

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/user?user_id=%v&email=%v", userUUID.String(), email)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Equal(t,
			"{\"error\":\"cannot use both user_id and email for querying\"}",
			recorder.Body.String(),
		)
	})
	t.Run("gets user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

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
		eur := ExpectedUserResult(cur, ti, tn)
		{
			store.
				EXPECT().
				GetUserByID(gomock.Any(), gomock.Eq(userUUID)).
				Times(1).
				Return(*cur, nil)
		}

		//postReq, err := json.Marshal(uur)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/user?user_id=%v", userUUID)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		checkResponseBody(t, eur, recorder.Body.Bytes())
	})
	t.Run("gets user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		userUUID, err := uuid.NewUUID()
		require.NoError(t, err)
		email := "some-email@test.com"
		tn := time.Now()
		ti := int64(4)

		cur := createUserResult(&user.CreateUserParams{
			UserID:    userUUID,
			Email:     email,
			FirstName: "some",
			LastName:  "email",
		}, ti, tn)
		eur := ExpectedUserResult(cur, ti, tn)
		{
			store.
				EXPECT().
				GetUserByEmail(gomock.Any(), gomock.Eq(email)).
				Times(1).
				Return(*cur, nil)
		}

		//postReq, err := json.Marshal(uur)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/user?email=%v", email)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		t.Log(recorder)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		checkResponseBody(t, eur, recorder.Body.Bytes())
	})
}

func TestServer_DeleteUser(t *testing.T) {
	t.Run("fails to delete user - both user_id and email provided", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		userUUID, err := uuid.NewUUID()
		require.NoError(t, err)
		email := "some/email@test.com"

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?user_id=%v&email=%v", userUUID.String(), email)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Equal(t,
			"{\"error\":\"cannot use both user_id and email for querying\"}",
			recorder.Body.String(),
		)
	})
	t.Run("fails to delete user - neither user_id and email provided", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users")
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Equal(t,
			"{\"error\":\"invalid query parameter for a single user\"}",
			recorder.Body.String(),
		)
	})
	t.Run("gets user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		userUUID, err := uuid.NewUUID()
		require.NoError(t, err)

		{
			store.
				EXPECT().
				DeleteUserByID(gomock.Any(), gomock.Eq(userUUID)).
				Times(1).
				Return(nil)
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?user_id=%v", userUUID)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		require.Equal(t,
			fmt.Sprintf(
				"{\"message\":\"user with uuid: %v - deleted successfully\"}",
				userUUID.String(),
			),
			recorder.Body.String(),
		)
	})
	t.Run("gets user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		email := "some-email@test.com"

		{
			store.
				EXPECT().
				DeleteUserByEmail(gomock.Any(), gomock.Eq(email)).
				Times(1).
				Return(nil)
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?email=%v", email)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		require.Equal(t,
			fmt.Sprintf(
				"{\"message\":\"user with email: %v - deleted successfully\"}",
				email,
			),
			recorder.Body.String(),
		)
	})
}

func TestServer_ListUsers(t *testing.T) {
	t.Run("fails to get users - both page_id and page_size must be provided", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		pqp := userModel.PaginationQueryParams{
			PageID:   1,
			PageSize: 5,
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?page_id=%v", pqp.PageID)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Equal(t,
			"{\"error\":\"cannot provide only one of the pagination elements\"}",
			recorder.Body.String(),
		)
	})
	t.Run("fails to get users - both page_id and page_size must be provided", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		pqp := userModel.PaginationQueryParams{
			PageID:   1,
			PageSize: 5,
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?page_size=%v", pqp.PageSize)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Equal(t,
			"{\"error\":\"cannot provide only one of the pagination elements\"}",
			recorder.Body.String(),
		)
	})
	t.Run("get users without pagination", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		xu, xur := createRandomUserList(t, 5)

		{
			store.
				EXPECT().
				ListUsers(gomock.Any()).
				Times(1).
				Return(xu, nil)
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users")
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		//checkResponseBody(t, xur, recorder.Body.Bytes())
		var userResponseList []userModel.UserResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &userResponseList)
		require.NoError(t, err)
		require.Equal(t, len(xur), len(userResponseList))
		for idx, userResponse := range userResponseList {
			bur, err := json.Marshal(userResponse)
			require.NoError(t, err)
			checkResponseBody(t, &xur[idx], bur)
		}
	})
	t.Run("get users with pagination", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdbcompany.NewMockStore(ctrl)
		server := userFactory.NewUserFactory(store)

		xu, xur := createRandomUserList(t, 5)

		pqp := userModel.PaginationQueryParams{
			PageID:   1,
			PageSize: 5,
		}

		pagParams := user.ListPaginatedUsersParams{
			Limit:  pqp.PageSize,
			Offset: (pqp.PageID - 1) * pqp.PageSize,
		}

		{
			store.
				EXPECT().
				ListPaginatedUsers(gomock.Any(), gomock.Eq(pagParams)).
				Times(1).
				Return(xu, nil)
		}

		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/users?page_size=%v&page_id=%v", pqp.PageSize, pqp.PageID)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
		server.Router.ServeHTTP(recorder, req)

		// check response
		require.Equal(t, http.StatusOK, recorder.Code)
		//checkResponseBody(t, xur, recorder.Body.Bytes())
		var userResponseList []userModel.UserResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &userResponseList)
		require.NoError(t, err)
		require.Equal(t, len(xur), len(userResponseList))
		for idx, userResponse := range userResponseList {
			bur, err := json.Marshal(userResponse)
			require.NoError(t, err)
			checkResponseBody(t, &xur[idx], bur)
		}
	})
}

func checkResponseBody(
	t *testing.T,
	expectedBody *userModel.UserResponse,
	body []byte,
) {
	var ur userModel.UserResponse
	err := json.Unmarshal(body, &ur)
	require.NoError(t, err)

	require.Equal(t, expectedBody.TableID, ur.TableID)
	require.Equal(t, expectedBody.Email, ur.Email)
	require.Equal(t, expectedBody.FirstName, ur.FirstName)
	require.Equal(t, expectedBody.LastName, ur.LastName)
	require.Truef(
		t,
		expectedBody.CreatedAt.Equal(ur.CreatedAt),
		fmt.Sprintf(
			"different timestamptz -> expected: %v | actual: %v\n",
			expectedBody.CreatedAt, ur.CreatedAt,
		),
	)
}
