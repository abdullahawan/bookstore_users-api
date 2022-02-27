package services

import (
	"bookstore_user-api/domain/users"
	"bookstore_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
