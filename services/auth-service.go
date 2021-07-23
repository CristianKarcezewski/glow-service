package services

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	IAuthService interface {
		Login(email, password string) (string, error)
	}
	authService struct{}
)

func NewAuthService() IAuthService {
	return &authService{}
}

func (auth *authService) Login(email, password string) (string, error) {

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

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
