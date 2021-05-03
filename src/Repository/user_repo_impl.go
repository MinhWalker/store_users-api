package user

import (
	"errors"
	"fmt"
	"github.com/MinhWalker/store_users-api/src/datasources/mysql/users_db"
	"github.com/MinhWalker/store_users-api/src/domain/users"
	"github.com/MinhWalker/store_users-api/src/logger"
	"github.com/MinhWalker/store_users-api/src/utils/mysql_utils"
	"github.com/MinhWalker/store_utils-go/rest_errors"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func NewUserRepository() UserRepository {
	return &usersRepository{}
}

type UserRepository interface {
	Get(users.User) rest_errors.RestErr
	Save(users.User) rest_errors.RestErr
	Update(users.User) rest_errors.RestErr
	Delete(users.User) rest_errors.RestErr
	FindByStatus(string) ([]users.User, rest_errors.RestErr)
	FindByEmailAndPassword(users.User) rest_errors.RestErr
}

type usersRepository struct {
}

func (u *usersRepository) Get(user users.User) rest_errors.RestErr {
	statement, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare get user", errors.New("database error"))
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	return nil
}

func (u *usersRepository) Save(user users.User) rest_errors.RestErr {
	statement, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare save user", errors.New("database error"))
	}
	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when trying to get last insert id after creating a new user", errors.New("database error"))
	}
	user.Id = userId
	return nil
}

func (u *usersRepository) Update(user users.User) rest_errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare update user", errors.New("database error"))
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}

	return nil
}

func (u *usersRepository) Delete(user users.User) rest_errors.RestErr {
	statement, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare delete user statement", errors.New("database error"))
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}

	return nil
}

func (u *usersRepository) FindByStatus(status string) ([]users.User, rest_errors.RestErr) {
	statement, err := users_db.Client.Prepare(queryFindByStatus)

	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to prepare find users by status statement", errors.New("database error"))
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		logger.Error("error when trying to prepare find users by status", err)
		return nil, rest_errors.NewInternalServerError("error when trying to prepare find users by status", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]users.User, 0)
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when scan user row into user struct", errors.New("database error"))
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (u *usersRepository) FindByEmailAndPassword(user users.User) rest_errors.RestErr {
	statement, err := users_db.Client.Prepare(queryFindByEmailAndPassword)

	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare get user by email and password statement", errors.New("database error"))
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email, user.Password, users.StatusActive)

	//TODO: scan matched row into the values pointed at by dest
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user by email and password", errors.New("database error"))
	}

	return nil
}
