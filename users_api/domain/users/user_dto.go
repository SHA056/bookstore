package users

import (
	"strings"

	"github.com/SHA056/bookstore/users_api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
