package services

import (
	"errors"
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/repository"

	"golang.org/x/crypto/bcrypt"
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
	log.AddStep("UserService-Register")

	log.AddInfo("Generating default user data")
	user.CreatedAt = functions.DateToString()
	user.LastLogin = functions.DateToString()
	user.Active = true
	user.UserGroupId = 1

	log.AddInfo("Encrypting password")
	daoUser := dao.NewDAOUser(user)

	newUser, userErr := us.userRepository.Insert(log, daoUser)
	if userErr != nil {
		return nil, userErr
	}

	hash, hashErr := us.hashPassword(&user.Password)
	if hashErr != nil {
		return nil, hashErr
	}

	hash.UserId = newUser.UserId
	saveHashErr := us.hashRepository.Insert(log, hash)
	if saveHashErr != nil {
		// TODO: DELETE NEW USER FROM DATABASE
		return nil, saveHashErr
	}

	return newUser.ToModel(), nil
}

func (us *userService) FindById(log *models.StackLog, userId *int64) (*models.User, error) {
	log.AddStep("UserService-FindById")

	us.userRepository.FindById(log, userId)
	return nil, nil
}

func (us *userService) VerifyUser(log *models.StackLog, email, password *string) (*models.User, error) {
	log.AddStep("UserService-VerifyUser")

	user, findErr := us.userRepository.Select(log, "email", email)
	if findErr != nil {
		return nil, findErr
	}

	databaseHash, databaseHashErr := us.hashRepository.Select(log, "user_id", user.UserId)
	if databaseHashErr != nil {
		return nil, databaseHashErr
	}

	if !us.checkPasswordHash(password, &databaseHash.Hash) {
		return nil, errors.New("user not recgonized")
	}
	return user.ToModel(), nil
}

func (us *userService) hashPassword(password *string) (*dao.Hash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &dao.Hash{Hash: string(bytes)}, nil
}

func (us *userService) checkPasswordHash(password *string, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
