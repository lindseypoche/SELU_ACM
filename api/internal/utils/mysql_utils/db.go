package mysql_utils

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
)

const (
	// ErrorNoRows ...
	ErrorNoRows = "no rows in result set"
)

// ParseError returns a rest error.
// This function should be used for all errors coming from a mysql database.
func ParseError(err error) rest.Err {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest.NewNotFoundError("no record matching given id")
		}
		return rest.NewInternalServerError("errors parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return rest.NewBadRequestError("invalid data")
	}
	return rest.NewInternalServerError("error processing request", errors.New("database error"))
}
