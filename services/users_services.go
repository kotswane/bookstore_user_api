package services

import (
	"github.com/kotswane/bookstore_user_api/domain/users"
	"github.com/kotswane/bookstore_user_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
	} else {
		current.Email = user.Email
		current.FirstName = user.FirstName
		current.LastName = user.LastName
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userid int64) *errors.RestErr {
	user := &users.User{Id: userid}
	return user.Delete()
}

func FindByStatus(status string) ([]users.User, *errors.RestErr) {
	return nil, nil
}
