package Repository

import (
	"github.com/MinhWalker/store_users-api/domain/users"
	"github.com/MinhWalker/store_users-api/utils/errors"
)

type UserRepository interface {
	Get() *errors.RestErr
	Save() *errors.RestErr
	Update() *errors.RestErr
	Delete() *errors.RestErr
	FindByStatus(status string) ([]users.User, *errors.RestErr)
	FindByEmailAndPassword() *errors.RestErr
}


