package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/kotswane/bookstore_user_api/utils/errors"
)

const (
	noRowFound = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowFound) {
			return errors.NewNotFoundRequestError("No matching record")
		}
		return errors.NewInternalServerError("Error parsing database record")
	}

	if sqlErr.Number == 1062 {
		return errors.NewBadRequestError("duplicate record")
	}

	return errors.NewInternalServerError("Error processing the request")
}
