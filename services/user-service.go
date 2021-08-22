package services

import (
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/repository"
)

type (
	IUserService interface {
		Register(log *models.StackLog, user *models.User) (*models.User, error)
		FindById(log *models.StackLog, userId *int64) (*models.User, error)
	}

	userService struct {
		userRepository repository.IUserRepository
	}
)

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{userRepository}
}

func (us *userService) Register(log *models.StackLog, user *models.User) (*models.User, error) {
	user.CreatedAt = functions.DateToString()
	user.Active = true
	user.UserGroupId = 1
	us.userRepository.Insert(user.AdaptToDAO())
	return nil, nil
}

func (us *userService) FindById(log *models.StackLog, userId *int64) (*models.User, error) {
	us.userRepository.FindById(userId)
	return nil, nil
}
