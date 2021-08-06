package repository

import (
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryUserTable = "users"
)

type (
	IUserRepository interface {
		Insert(user *dao.UserDao) error
	}
	userRepository struct {
		database server.IDatabaseHandler
	}
)

func NewUserRepository(database server.IDatabaseHandler) IUserRepository {
	return &userRepository{database}
}

func (ur *userRepository) Insert(user *dao.UserDao) error {
	return ur.database.Insert(repositoryUserTable, user)
}
