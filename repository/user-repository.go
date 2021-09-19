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
		Insert(log *models.StackLog, user *dao.User) (*dao.User, error)
		Select(log *models.StackLog, key string, value interface{}) (*dao.User, error)
		Update(log *models.StackLog, user *dao.User) (*dao.User, error)
	}
	userRepository struct {
		database server.IDatabaseHandler
	}
)

func NewUserRepository(database server.IDatabaseHandler) IUserRepository {
	return &userRepository{database}
}

func (ur *userRepository) Insert(log *models.StackLog, user *dao.User) (*dao.User, error) {
	log.AddStep("UserRepository-Insert")

	log.AddInfo("Saving user")
	err := ur.database.Insert(repositoryUserTable, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Select(log *models.StackLog, key string, value interface{}) (*dao.User, error) {
	log.AddStep("UserRepository-Select")

	var user dao.User

	log.AddInfo(fmt.Sprintf("Finding user by %s", key))
	selectErr := ur.database.Select(repositoryUserTable, &user, key, value)
	if selectErr != nil {
		return nil, selectErr
	}
	return &user, nil
}

func (ur *userRepository) Update(log *models.StackLog, user *dao.User) (*dao.User, error) {
	log.AddStep("UserRepository-Update")

	err := ur.database.Update(repositoryUserTable, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
