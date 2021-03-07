package mysql_utils

import (
	"atnlie/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.CustomNotFoundError("no record matching gived id")
		}
		return errors.CustomInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.CustomBadRequestError("invalid data")
	case 1364:
		return errors.CustomBadRequestError("field does not have a default value")
	case 1064:
		return errors.CustomBadRequestError("SQL Syntax is error")
	}
	return errors.CustomInternalServerError("error parsing process")
}
