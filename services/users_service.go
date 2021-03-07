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

func UpdateUser(isPartial bool, users users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(users.Id)
	if err != nil {
		return nil, err
	}

	if err := users.UserValidation(); err != nil {
		return nil, err
	}

	if isPartial {
		if users.FirstName != "" {
			currentUser.FirstName = users.FirstName
		}

		if users.LastName != "" {
			currentUser.LastName = users.LastName
		}

		if users.Email != "" {
			currentUser.Email = users.Email
		}
	} else {
		currentUser.FirstName = users.FirstName
		currentUser.LastName = users.LastName
		currentUser.Email = users.Email
	}
	//always update date created otherwise make new field to store update date
	currentUser.DateCreated = date_utils.GetNowString()

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	currentUser := &users.User{Id: userId}
	return currentUser.Delete()
}

func SearchUser(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}