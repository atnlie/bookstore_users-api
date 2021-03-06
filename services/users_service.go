package services

import (
	"atnlie/domain/users"
	"atnlie/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	/*
	//call function
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	return &user, nil
	*/

	//call methods
	if err := user.UserValidation(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}