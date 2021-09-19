package presenters

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IAuthPresenter interface {
		RefreshToken() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	authPresenter struct {
		errorMessageData *models.ServerErrorMessages
		authService      services.IAuthService
	}
)

func NewAuthPresenter(errorMessageData *models.ServerErrorMessages, authService services.IAuthService) IAuthPresenter {
	return &authPresenter{errorMessageData, authService}
}

func (ac *authPresenter) RefreshToken() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "")
	}
}

// Router is a function that returns a router of authPresenter
func (ac *authPresenter) Router(echo *echo.Echo) {
	router := routers.AuthRouter{
		Echo:         echo,
		RefreshToken: ac.RefreshToken(),
	}

	router.Wire()
}
