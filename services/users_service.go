package services

import "atnlie/domain/users"

func CreateUser(user users.User) (*users.User, error){
	return &user, nil
}