package services

import (
	"github.com/MinhWalker/store_users-api/domain/users"
	"github.com/MinhWalker/store_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr)  {
	return nil, nil

	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}