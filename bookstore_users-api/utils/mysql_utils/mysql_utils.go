package mysql_utils

import (
	"github.com/esboych/microservices-book/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	log.Println("Mysql error:", err)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequsetError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		//debug.PrintStack()
		return errors.NewBadRequsetError("invalid data")

	}
	return errors.NewInternalServerError("error processing request")
}
