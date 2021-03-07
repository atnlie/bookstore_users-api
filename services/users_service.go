package services

import (
	"atnlie/domain/users"
	"atnlie/utils/crypto_utils"
	"atnlie/utils/date_utils"
	"atnlie/utils/errors"
)

var (
	UserService UsersServiceInterface = &userService{}
)

type userService struct{}

type UsersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
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

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUser(isPartial bool, users users.User) (*users.User, *errors.RestErr) {
	currentUser, err := UserService.GetUser(users.Id)
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

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	currentUser := &users.User{Id: userId}
	return currentUser.Delete()
}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
