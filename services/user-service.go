package services

import (
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/repository"
)

type (
	IUserService interface {
		Register(log *models.StackLog, user *models.User) (*models.User, error)
		FindById(log *models.StackLog, userId *int64) (*models.User, error)
		VerifyUser(log *models.StackLog, email, password *string) (*models.User, error)
	}

	userService struct {
		userRepository repository.IUserRepository
		hashRepository repository.IHashRepository
	}
)

func NewUserService(userRepository repository.IUserRepository, hashRepository repository.IHashRepository) IUserService {
	return &userService{userRepository, hashRepository}
}

func (us *userService) Register(log *models.StackLog, user *models.User) (*models.User, error) {

	user.CreatedAt = functions.DateToString()
	user.LastLogin = functions.DateToString()
	user.Active = true
	user.UserGroupId = 1

	daoUser, hash, hashErr := dao.NewDAOUser(user)
	if hashErr != nil {
		return nil, hashErr
	}

	newUser, userErr := us.userRepository.Insert(daoUser)
	if userErr != nil {
		return nil, userErr
	}

	hash.UserId = newUser.UserId
	saveHashErr := us.hashRepository.Insert(hash)
	if saveHashErr != nil {
		// TODO: DELETE NEW USER FROM DATABASE
		return nil, saveHashErr
	}

	return newUser.ToModel(), nil
}

func (us *userService) FindById(log *models.StackLog, userId *int64) (*models.User, error) {
	us.userRepository.FindById(userId)
	return nil, nil
}

func (us *userService) VerifyUser(log *models.StackLog, email, password *string) (*models.User, error) {
	return nil, nil
}
