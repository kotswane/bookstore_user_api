package users

import (
	"strings"

	"github.com/kotswane/bookstore_user_api/utils/errors"
)

type User struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Date_Created string `json:"date_created"`
	Status       string `json:"status"`
	Password     string `json:"-"`
}

func (users *User) Validate() *errors.RestErr {
	users.Email = strings.TrimSpace(strings.ToLower(users.Email))
	if users.Email == "" {
		return errors.NewBadRequestError("invalid email")
	}
	return nil
}
