package services

import (
	"context"
	"errors"
	"glow-service/models"

	"firebase.google.com/go/auth"
)

type (
	IAuthService interface {
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

	log.SetUser(user.Email)
	return &user, nil
}

func (auth *authService) GenerateToken(log *models.StackLog, user *models.User) (*models.Auth, error) {
	log.AddStep("AuthService-GenerateToken")
	log.SetUser(user.Email)

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
