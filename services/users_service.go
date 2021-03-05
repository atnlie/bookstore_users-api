package services

import (
	"atnlie/domain/users"
	"atnlie/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr){
	return &user, nil
}