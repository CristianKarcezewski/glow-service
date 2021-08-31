package services

import (
	"errors"
	"glow-service/common/functions"
	"glow-service/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	tokenSecret = "churiChurinFunFlais"
)

type (
	IAuthService interface {
		Login(log *models.StackLog, email, password *string) (*models.Auth, error)
		Register(log *models.StackLog, user *models.User) (*models.Auth, error)
		VerifyToken(log *models.StackLog, token string) (*models.User, error)
	}
	authService struct {
		userService IUserService
	}
)

func NewAuthService(userService IUserService) IAuthService {
	return &authService{userService}
}

func (auth *authService) Login(log *models.StackLog, email, password *string) (*models.Auth, error) {
	log.AddStep("AuthService-Login")

	user, userErr := auth.userService.VerifyUser(log, email, password)
	if userErr != nil {
		return nil, userErr
	}

	log.AddInfo("Generating token")
	token, tokenErr := auth.generateToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &models.Auth{Authorization: token}, nil
}

func (auth *authService) Register(log *models.StackLog, user *models.User) (*models.Auth, error) {
	log.AddStep("AuthService-Register")

	user, userErr := auth.userService.Register(log, user)
	if userErr != nil {
		return nil, userErr
	}

	log.AddInfo("Generating token")
	token, tokenErr := auth.generateToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &models.Auth{Authorization: token}, nil

}

func (auth *authService) VerifyToken(log *models.StackLog, tokenStr string) (*models.User, error) {

	log.AddStep("AuthService-ValidateToken")

	var user models.User
	hmacSecret := []byte(tokenSecret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		user.UserId = claims["userId"].(int64)
		user.UserGroupId = claims["userGroupId"].(int64)
		user.UserName = claims["name"].(string)
		user.Email = claims["email"].(string)

		exp := claims["exp"].(string)
		dt, _ := functions.StringToDate(exp)
		if auth.compareTokenDate(dt) {
			log.SetUser(user.Email)
			return &user, nil
		}

	}
	return nil, errors.New("Invalid token")
}

func (auth *authService) compareTokenDate(date time.Time) bool {
	currentTime := time.Now()
	invalidTime := currentTime.Add(-time.Minute * 60)

	diff := invalidTime.Before(date)
	return diff
}

func (auth *authService) generateToken(user *models.User) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.UserId
	claims["userGroupId"] = user.UserGroupId
	claims["name"] = user.UserName
	claims["email"] = user.Email
	claims["exp"] = functions.DateToString()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(tokenSecret))
}
