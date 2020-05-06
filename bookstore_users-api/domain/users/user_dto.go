package users

import (
	"fmt"
	"github.com/esboych/microservices-book/bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)
	fmt.Println("pass:", user.Password)

	if user.Email == "" {
		//fmt.Println("user email validation ERROR id dto")
		return errors.NewBadRequsetError("email address cannot be empty")
	}

	if user.Password == "" {
		//fmt.Println("user password validation ERROR id dto")
		return errors.NewBadRequsetError("invalid password: should not be empty")
	}


	return nil
}
