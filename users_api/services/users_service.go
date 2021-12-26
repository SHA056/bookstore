package services

import (
	"net/http"

	"github.com/SHA056/bookstore/users_api/domain/users"
	"github.com/SHA056/bookstore/users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
