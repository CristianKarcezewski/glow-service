package repository

import (
	"fmt"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryUserTable = "users"
)

type (
	IUserRepository interface {
		Insert(user *dao.User) error
		FindById(userId *int64) (*dao.User, error)
	}
	userRepository struct {
		database server.IDatabaseHandler
	}
)

func NewUserRepository(database server.IDatabaseHandler) IUserRepository {
	return &userRepository{database}
}

func (ur *userRepository) Insert(user *dao.User) error {
	return ur.database.Insert(repositoryUserTable, user)
}

func (ur *userRepository) FindById(userId *int64) (*dao.User, error) {
	var user models.User
	us, err := ur.database.FindById(repositoryUserTable, userId, &user)
	if err != nil {
		return nil, err
	}
	if us != nil {
		fmt.Println(us)
	}
	return nil, nil
}
