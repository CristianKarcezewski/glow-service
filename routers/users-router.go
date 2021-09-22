package routers

import "github.com/labstack/echo"

const (
	userLoginPath    = "/users/login"
	userGetPath      = "/users/:userId"
	userRegisterPath = "/users"
	userUpdatePath   = "/users"
)

type (
	UsersRouter struct {
		Echo     *echo.Echo
		Login    echo.HandlerFunc
		GetById  echo.HandlerFunc
		Register echo.HandlerFunc
		Update   echo.HandlerFunc
	}
)

func (ur *UsersRouter) Wire() {
	ur.Echo.POST(userLoginPath, ur.Login)
	ur.Echo.GET(userGetPath, ur.GetById)
	ur.Echo.POST(userRegisterPath, ur.Register)
	ur.Echo.PUT(refreshTokenPath, ur.Update)
}
