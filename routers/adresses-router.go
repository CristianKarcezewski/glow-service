package routers

import "github.com/labstack/echo"

const (
	registerAddressPath = "/addresses/register"
)

type (
	AddressesRouter struct {
		Echo     *echo.Echo
		Register echo.HandlerFunc
	}
)

func (ar *AddressesRouter) Wire() {
	ar.Echo.POST(registerAddressPath, ar.Register)
}
