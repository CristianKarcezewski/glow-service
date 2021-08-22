package services

import (
	"glow-service/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	IAuthService interface {
		Login(log *models.StackLog, email, password *string) (string, error)
		Register(log *models.StackLog, user *models.User) (*string, error)
	}
	authService struct {
		userService IUserService
	}
)

func NewAuthService(userService IUserService) IAuthService {
	return &authService{userService}
}

func (auth *authService) Login(log *models.StackLog, email, password *string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

func (auth *authService) Register(log *models.StackLog, user *models.User) (*string, error) {
	_, userErr := auth.userService.Register(log, user)
	if userErr != nil {
		return nil, userErr
	}
	return nil, nil
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
