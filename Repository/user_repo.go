package Repository

import (
	"github.com/MinhWalker/store_users-api/domain/users"
	"github.com/MinhWalker/store_utils-go/rest_errors"
)

type UserRepository interface {
	Get() *rest_errors.RestErr
	Save() *rest_errors.RestErr
	Update() *rest_errors.RestErr
	Delete() *rest_errors.RestErr
	FindByStatus(status string) ([]users.User, *rest_errors.RestErr)
	FindByEmailAndPassword() *rest_errors.RestErr
}


