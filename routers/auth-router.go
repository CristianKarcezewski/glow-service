package routers

import "github.com/labstack/echo"

const (
	loginPath        = "/login"
	refreshTokenPath = "/refresh"
	registerPath     = "/user/register"
)

type (
	AuthRouter struct {
		Echo         *echo.Echo
		Login        echo.HandlerFunc
		RefreshToken echo.HandlerFunc
		Register     echo.HandlerFunc
	}
)

// Wire is the function that serve the router
func (ar *AuthRouter) Wire() {
	ar.Echo.POST(loginPath, ar.Login)
	ar.Echo.GET(refreshTokenPath, ar.RefreshToken)
	ar.Echo.POST(registerPath, ar.Register)
}
