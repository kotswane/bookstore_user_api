package services

import (
	"github.com/kotswane/bookstore_user_api/domain/user"
	"github.com/kotswane/bookstore_user_api/utils/errors"
)

func CreateUser(user user.User) (*user.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*user.User, *errors.RestErr) {
	result := &user.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
