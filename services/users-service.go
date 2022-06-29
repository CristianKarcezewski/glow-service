package services

import (
	"context"
	"errors"
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/models/dto"
	"glow-service/repository"

	"firebase.google.com/go/auth"
	"golang.org/x/crypto/bcrypt"
)

type (
	IUsersService interface {
		Login(log *models.StackLog, auth *dto.AuthData) (*models.Auth, error)
		Register(log *models.StackLog, user *models.User) (*models.Auth, error)
		GetById(log *models.StackLog, userId int64) (*models.User, error)
		VerifyUser(log *models.StackLog, email, password *string) (*models.User, error)
		Update(log *models.StackLog, user *models.User) (*models.User, error)
	}

	usersService struct {
		userRepository repository.IUserRepository
		hashRepository repository.IHashRepository
		FirebaseClient *auth.Client
		authService    IAuthService
	}
)

func NewUserService(firebaseClient *auth.Client, userRepository repository.IUserRepository, hashRepository repository.IHashRepository, authService IAuthService) IUsersService {
	return &usersService{
		userRepository,
		hashRepository,
		firebaseClient,
		authService,
	}
}

func (us *usersService) Login(log *models.StackLog, auth *dto.AuthData) (*models.Auth, error) {
	log.AddStep("UserService-Login")
	user, userErr := us.VerifyUser(log, &auth.Email, &auth.Password)
	if userErr != nil {
		return nil, userErr
	}

	authData, authErr := us.authService.GenerateToken(log, user)
	if authErr != nil {
		return nil, authErr
	}
	return authData, nil
}

func (us *usersService) Register(log *models.StackLog, user *models.User) (*models.Auth, error) {
	log.AddStep("UserService-Register")

	log.AddInfo("Generating default user data")
	user.CreatedAt = functions.DateToString(nil)
	user.Active = true
	user.UserGroupId = 1

	//register user in firebase
	uid, firErr := us.createFirebaseUser(log, user)
	if firErr != nil {
		return nil, firErr
	}
	user.Uid = *uid

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

	auth, authErr := us.authService.GenerateToken(log, newUser.ToModel())
	if authErr != nil {
		return nil, authErr
	}

	return auth, nil
}

func (us *usersService) GetById(log *models.StackLog, userId int64) (*models.User, error) {
	log.AddStep("UserService-FindById")

	user, err := us.userRepository.Select(log, "id", userId)
	if err != nil {
		return nil, err
	}
	return user.ToModel(), nil
}

func (us *usersService) VerifyUser(log *models.StackLog, email, password *string) (*models.User, error) {
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

func (us *usersService) Update(log *models.StackLog, user *models.User) (*models.User, error) {
	log.AddStep("UserService-Update")

	dbUser, findErr := us.GetById(log, user.UserId)
	if findErr != nil {
		return nil, findErr
	}

	if user.UserName != "" {
		dbUser.UserName = user.UserName
	}
	if user.ImageUrl != dbUser.ImageUrl {
		dbUser.ImageUrl = user.ImageUrl
	}
	if user.UserGroupId > 0 && user.UserGroupId != dbUser.UserGroupId {
		dbUser.UserGroupId = user.UserGroupId
	}
	if user.Email != dbUser.Email {
		dbUser.Email = user.Email
	}
	if user.Phone != dbUser.Phone {
		dbUser.Phone = user.Phone
	}

	daoUser, err := us.userRepository.Update(log, dao.NewDAOUser(dbUser))
	if err != nil {
		return nil, err
	}
	return daoUser.ToModel(), nil
}

func (us *usersService) hashPassword(password *string) (*dao.Hash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &dao.Hash{Hash: string(bytes)}, nil
}

func (us *usersService) checkPasswordHash(password *string, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}

func (us *usersService) createFirebaseUser(log *models.StackLog, user *models.User) (*string, error) {
	log.AddInfo("Creating firebase user")

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.UserName).
		Disabled(false)

	u, err := us.FirebaseClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, errors.New("error creating firebase user")
	}
	return &u.UID, nil
}
