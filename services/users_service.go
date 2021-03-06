package services

import (
	"atnlie/domain/users"
	"atnlie/utils/date_utils"
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

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(users users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(users.Id)
	if err != nil {
		return nil, err
	}

	currentUser.FirstName = users.FirstName
	currentUser.LastName = users.LastName
	currentUser.Email = users.Email
	currentUser.DateCreated = date_utils.GetNowString()

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}
