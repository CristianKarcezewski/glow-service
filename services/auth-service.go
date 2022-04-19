package services

import (
	"context"
	"errors"
	"glow-service/models"

	"firebase.google.com/go/auth"
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
	token, err := auth.firebaseClient.VerifyIDToken(context.Background(), tokenStr)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims
	if userGroupId, ok := claims["admin"]; ok {
		if userGroupId.(bool) {
			return &models.User{
				UserGroupId: userGroupId.(int64),
			}, nil
		}

	}

	return nil, errors.New("invalid token")
}

func (auth *authService) GenerateToken(log *models.StackLog, user *models.User) (*models.Auth, error) {
	log.AddStep("AuthService-GenerateToken")

	// // Set claims
	claims := map[string]interface{}{
		"userGroupId": user.UserGroupId,
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
