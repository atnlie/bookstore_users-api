package users

import (
	"atnlie/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

const (
	StatusActive   = "active"
	StatusDeactive = "deactive"
)

//function
func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.CustomBadRequestError("Invalid email address")
	}
	return nil
}

//method
func (user *User) UserValidation() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.CustomBadRequestError("Invalid Email Address")
	}

	user.Password = strings.TrimSpace(user.Password)
	println("pass ", user.Password)
	if user.Password == "" {
		return errors.CustomBadRequestError("Invalid Password")
	}
	return nil
}
