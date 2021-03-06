package mysql_utils

import (
	"atnlie/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr  {
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
	}
	return errors.CustomInternalServerError("error parsing process")
}
