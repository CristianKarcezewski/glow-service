package services

import (
	"context"
	"errors"
	"glow-service/models"
	"time"

	"firebase.google.com/go/auth"
)

const (
	tokenSecret = "senhaSuperSecreta"
)

type (
	IAuthService interface {
		// Login(log *models.StackLog, email, password *string) (*models.Auth, error)
		GenerateToken(log *models.StackLog, user *models.User) (*models.Auth, error)
		VerifyToken(log *models.StackLog, token string) (*models.User, error)
	}

	authService struct {
		firebaseClient *auth.Client
	}
)

func NewAuthService(firebaseClient *auth.Client) IAuthService {
	return &authService{firebaseClient}
}

func (auth *authService) VerifyToken(log *models.StackLog, tokenStr string) (*models.User, error) {

	log.AddStep("AuthService-ValidateToken")

	// var user models.User
	// hmacSecret := []byte(tokenSecret)
	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	// 	// check token signing method etc
	// 	return hmacSecret, nil
	// })
	// if err != nil {
	// 	return nil, errors.New("invalid token")
	// }
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	// user.UserId = claims["userId"].(float64)
	// 	// user.UserGroupId = claims["userGroupId"].(float64)
	// 	user.UserId = int64(claims["userId"].(float64))
	// 	user.UserGroupId = int64(claims["userGroupId"].(float64))
	// 	user.UserName = claims["name"].(string)
	// 	user.Email = claims["email"].(string)
	// 	exp := claims["exp"].(string)
	// 	dt, _ := functions.StringToDate(exp)
	// 	if !auth.compareTokenDate(dt) {
	// 		log.SetUser(user.Email)
	// 		return &user, nil
	// 	}
	// }
	// return nil, errors.New("invalid token")

	token, err := auth.firebaseClient.VerifyIDToken(context.Background(), tokenStr)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims
	user := models.User{
		UserGroupId: int64(claims["userGroupId"].(float64)),
		UserId:      int64(claims["userId"].(float64)),
		UserName:    claims["name"].(string),
		Email:       claims["email"].(string),
	}
	// if userGroupId, ok := claims["userGroupId"]; ok {
	// 	if userGroupId.(bool) {
	// 		user.UserGroupId = userGroupId.(int64)
	// 	} else {
	// 		return nil, errors.New("invalid token")
	// 	}
	// }

	// if userId, ok := claims["userId"]; ok {
	// 	if userId.(bool) {
	// 		user.UserId = userId.(int64)
	// 	} else {
	// 		return nil, errors.New("invalid token")
	// 	}
	// }

	log.SetUser(user.Email)
	return &user, nil
}

func (auth *authService) GenerateToken(log *models.StackLog, user *models.User) (*models.Auth, error) {
	log.AddStep("AuthService-GenerateToken")

	// dateTime := functions.DateToString()

	// // Create token
	// token := jwt.New(jwt.SigningMethodHS256)

	// // Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["userId"] = user.UserId
	// claims["userGroupId"] = user.UserGroupId
	// claims["name"] = user.UserName
	// claims["email"] = user.Email
	// claims["exp"] = dateTime

	// tokenStr, err := token.SignedString([]byte(tokenSecret))
	// if err != nil {
	// 	return nil, errors.New("error creating user token")
	// }

	// return &models.Auth{
	// 	UserId:        user.UserId,
	// 	UserGroupId:   user.UserGroupId,
	// 	Expirate:      dateTime,
	// 	Authorization: tokenStr,
	// }, nil

	// // Set claims
	claims := map[string]interface{}{
		"userGroupId": user.UserGroupId,
		"userId":      user.UserId,
		"name":        user.UserName,
	}

	token, err := auth.firebaseClient.CustomTokenWithClaims(context.Background(), user.Uid, claims)
	if err != nil {
		return nil, errors.New("error creating user token")
	}
	return &models.Auth{
		UserId:        user.UserId,
		UserGroupId:   user.UserGroupId,
		Authorization: token,
	}, nil
}

func (auth *authService) compareTokenDate(date time.Time) bool {
	currentTime := time.Now()
	invalidTime := currentTime.Add(-time.Minute * 60)
	diff := invalidTime.Before(date)
	return diff
}
