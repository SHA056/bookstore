package users

import (
	"fmt"

	"github.com/SHA056/bookstore/users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Firstname = result.Firstname
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user

	return nil
}
