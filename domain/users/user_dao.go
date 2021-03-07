package users

import (
	"atnlie/datasources/mysql/users_db"
	"atnlie/utils/errors"
	"atnlie/utils/mysql_utils"
	"fmt"
)

//user only when you use array mode only
//var (
//	userDB = make(map[int64]*User)
//)

const (
	qryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created, status, password) " +
		"VALUES (?, ?, ?, ?, ?, ?);"
	qryGetUser      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	qryUpdateUser   = "UPDATE users SET first_name=?, last_name=?, email=?, date_created=? WHERE id=?;"
	qryDeleteUser   = "DELETE FROM users WHERE id=?;"
	qryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.ClientDb.Prepare(qryGetUser)
	if err != nil {
		return errors.CustomInternalServerError(err.Error())
	}
	//this for db mode
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status);
		err != nil {
		fmt.Println("err -> ", err)
		return mysql_utils.ParseError(err)
		//if strings.Contains(err.Error(), errorNoRows) {
		//	return errors.CustomNotFoundError(fmt.Sprintf("User: %d not found", user.Id))
		//}
		//
		//return errors.CustomInternalServerError(fmt.Sprintf("User: %d not found, %s", user.Id, err.Error()))
	}

	// if you want to use stmt.Query instead dont forget to close manually
	/*
		qResult, err := stmt.Query(user.Id)
		if err != nil {
			return errors.CustomInternalServerError(fmt.Sprintf("User: %d not found, %s", user.Id, err.Error()))
		}
		defer qResult.Close()
	*/

	return nil

	//this for array mode
	/*
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
	*/
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.ClientDb.Prepare(qryInsertUser)
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

	insertResult, saveErr := stmt.Exec(
		user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		fmt.Println("err ", saveErr)
		return mysql_utils.ParseError(saveErr)
		////add handle error using mysqlError
		//sqlErr, ok := saveErr.(*mysql.MySQLError)
		//if !ok {
		//	return errors.CustomInternalServerError(fmt.Sprintf("Error when trying to save user: %s", saveErr.Error()))
		//}
		//
		////use sqlErr detect
		//switch sqlErr.Number {
		//case 1062: //duplicate
		//	return errors.CustomBadRequestError(fmt.Sprintf("Email: %s already registered", user.Email))
		//}

		//instead of use manual check just use detect by mysql error code
		/*
			if strings.Contains(saveErr.Error(), errorEmailDuplicate) {
				return errors.CustomBadRequestError(fmt.Sprintf("Email: %s already registered", user.Email))
			}
		*/
		//return errors.CustomInternalServerError(fmt.Sprintf("Error when trying to save user: %s", saveErr.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.ClientDb.Prepare(qryUpdateUser)
	if err != nil {
		return errors.CustomInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.ClientDb.Prepare(qryDeleteUser)
	if err != nil {
		return errors.CustomInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.CustomInternalServerError(err.Error())
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.ClientDb.Prepare(qryFindByStatus)
	if err != nil {
		return nil, errors.CustomInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.CustomInternalServerError(err.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.CustomNotFoundError(fmt.Sprintf("no users matching status: %s", status))
	}
	return result, nil
}
