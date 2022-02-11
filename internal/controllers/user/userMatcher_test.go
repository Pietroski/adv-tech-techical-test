package userController_test

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	//"github.com/google/uuid"
	"github.com/golang/mock/gomock"

	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
)

func Eq(x interface{}) Matcher { return eqMatcher{x} }

type Matcher interface {
	// Matches returns whether x is a match.
	Matches(x interface{}) bool

	// String describes what the matcher matches.
	String() string
}

type eqMatcher struct {
	x interface{}
}

func (e eqMatcher) Matches(x interface{}) bool {
	return reflect.DeepEqual(e.x, x)
}

func (e eqMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.x)
}

type eqCreateUserParamsMatcher struct {
	arg user.CreateUserParams
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(user.CreateUserParams)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.arg.Email, arg.Email) &&
		reflect.DeepEqual(e.arg.FirstName, arg.FirstName) &&
		reflect.DeepEqual(e.arg.LastName, arg.LastName) &&
		!reflect.DeepEqual(arg.UserID, uuid.Nil)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("does not match arg %v", e.arg)
}

func EqCreateCompanyParams(arg user.CreateUserParams) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg}
}

type eqCreateUserResponseMatcher struct {
	arg user.Users
}

func (e eqCreateUserResponseMatcher) Matches(x interface{}) bool {
	arg, ok := x.(user.Users)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.arg.Email, arg.Email) &&
		reflect.DeepEqual(e.arg.FirstName, arg.FirstName) &&
		reflect.DeepEqual(e.arg.LastName, arg.LastName) // &&
	//!reflect.DeepEqual(arg.UserID, uuid.Nil) &&
	//!reflect.DeepEqual(arg.CreatedAt, time.Time{}) &&
	//!reflect.DeepEqual(arg.TableID, sql.NullInt64{})
}

func (e eqCreateUserResponseMatcher) String() string {
	return fmt.Sprintf("does not match arg %v", e.arg)
}

func EqCreateCompanyResponse(arg user.Users) gomock.Matcher {
	return eqCreateUserResponseMatcher{arg}
}
