package users

import (
	"atnlie/datasources/mysql/users_db"
	"atnlie/utils/date_utils"
	"atnlie/utils/errors"
	"fmt"
	"strings"
)

var (
	userDB = make(map[int64]*User)
)

const (
	QryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	QryGetUser    = "SELECT * FROM users WHERE id ="
	EmailDuplicate = "Duplicate"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.ClientDb.Ping(); err != nil {
		panic(err)
	}

	result := userDB[user.Id]
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
	stmt, err := users_db.ClientDb.Prepare(QryInsertUser)
	if err != nil {
		return errors.CustomInternalServerError(err.Error())
	}
	//this for db mode
	defer stmt.Close()
	//fast way with one line code
	/*
		result, err := users_db.ClientDb.Exec(qryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
		if err != nil {
			return errors.CustomInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
		}
	*/
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), EmailDuplicate) {
			return errors.CustomBadRequestError(fmt.Sprintf("Email: %s already registered", user.Email))
		}
		return errors.CustomInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.CustomInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil

	//this for array mode
	/*
		currentUser := userDB[user.Id]
		if currentUser != nil {
			if currentUser.Email == user.Email {
				return errors.CustomBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
			}
			return errors.CustomBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
		}

		user.DateCreated = date_utils.GetNowString()
		userDB[user.Id] = user
		return nil
	*/
}
