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
		FindById(log *models.StackLog, userId *int64) (*dao.User, error)
		Select(log *models.StackLog, key string, value interface{}) (*dao.User, error)
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

func (ur *userRepository) FindById(log *models.StackLog, userId *int64) (*dao.User, error) {
	log.AddStep("UserRepository-FindById")

	// var user models.User
	// us, err := ur.database.FindById(repositoryUserTable, userId, &user)
	// if err != nil {
	// 	return nil, err
	// }
	// if us != nil {
	// 	fmt.Println(us)
	// }
	return nil, nil
}
