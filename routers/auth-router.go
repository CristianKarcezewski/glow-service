package routers

import "github.com/labstack/echo"

const (
	loginPath        = "/login"
	refreshTokenPath = "/refresh"
)

type (
	AuthRouter struct {
		Echo         *echo.Echo
		RefreshToken echo.HandlerFunc
	}
)

// Wire is the function that serve the router
func (ar *AuthRouter) Wire() {
	ar.Echo.GET(refreshTokenPath, ar.RefreshToken)
}
