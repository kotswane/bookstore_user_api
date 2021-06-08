package user

import (
	"fmt"

	"github.com/kotswane/bookstore_user_api/utils/errors"
)

var (
	userDB = make(map[int16]*User)
)

func (user *User) Get() *errors.RestErr {

	result := userDB[int16(user.Id)]

	if result == nil {
		return errors.NewNotFoundRequestError(fmt.Sprintf("userId %d not found", user.Id))
	}

	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Date_Created = result.Date_Created
	user.Id = result.Id
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[int16(user.Id)]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("userId %d already exists", user.Id))
	}
	userDB[int16(user.Id)] = user
	return nil
}
