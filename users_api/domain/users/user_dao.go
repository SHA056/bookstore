package users

import (
	"fmt"

	"github.com/SHA056/bookstore/users_api/datasources/mysql/users_db"
	"github.com/SHA056/bookstore/users_api/utils/dates"
	"github.com/SHA056/bookstore/users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	user.DateCreated = dates.GetNowString()

	usersDB[user.Id] = user
	return nil
}
