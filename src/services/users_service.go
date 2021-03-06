package services

import (
<<<<<<< HEAD
<<<<<<< HEAD:src/services/users_service.go
	"github.com/MinhWalker/store_users-api/src/Repository/user"
||||||| merged common ancestors
<<<<<<<<< Temporary merge branch 1:src/services/users_service.go
=======
<<<<<<< HEAD:src/services/users_service.go
>>>>>>> f9578d851dcedaef23f32998606659f2268f00f3
	"github.com/MinhWalker/store_users-api/src/domain/users"
	"github.com/MinhWalker/store_users-api/src/utils/crypto_utils"
	"github.com/MinhWalker/store_users-api/src/utils/date_utils"
||||||| merged common ancestors:services/users_service.go
	"github.com/MinhWalker/store_users-api/domain/users"
	"github.com/MinhWalker/store_users-api/utils/crypto_utils"
	"github.com/MinhWalker/store_users-api/utils/date_utils"
<<<<<<< HEAD
=======
||||||| merged common ancestors
=========
	"github.com/MinhWalker/store_users-api/src/Repository/user"
=======
=======
	"github.com/MinhWalker/store_users-api/src/Repository/user"
>>>>>>> f9578d851dcedaef23f32998606659f2268f00f3
	"github.com/MinhWalker/store_users-api/src/domain/users"
	"github.com/MinhWalker/store_users-api/src/utils/crypto_utils"
	"github.com/MinhWalker/store_users-api/src/utils/date_utils"
<<<<<<< HEAD
>>>>>>> main:services/users_service.go
||||||| merged common ancestors
>>>>>>>>> Temporary merge branch 2:services/users_service.go
=======
>>>>>>> 44da521cf0b587bfbf3c25a1761f0c41c62de62d:services/users_service.go
>>>>>>> f9578d851dcedaef23f32998606659f2268f00f3
	"github.com/MinhWalker/store_utils-go/rest_errors"
)

type UsersService interface {
	GetUser(int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

type usersService struct{
	userRepository user.UserRepository
}

func NewUserService(userRepo user.UserRepository) UsersService {
	return &usersService{
		userRepository: userRepo,
	}
}

func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	user := users.User{Id: userId}
	if err := s.userRepository.Get(user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current := users.User{Id: user.Id}
	if err := s.userRepository.Get(current); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := s.userRepository.Update(current); err != nil {
		return nil, err
	}
	return &current, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	current := users.User{Id: userId}
	return s.userRepository.Delete(current)
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	return s.userRepository.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	user := users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := s.userRepository.FindByEmailAndPassword(user); err != nil {
		return nil, err
	}
	return &user, nil
}

