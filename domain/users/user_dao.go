package users

import (
	"atnlie/utils/errors"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	fmt.Println("Result ", result)
	if result == nil {
		return errors.CustomNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	currentUser := userDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.CustomBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.CustomBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
	}

	userDB[user.Id] = user
	return nil
}