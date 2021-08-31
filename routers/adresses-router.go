package routers

import "github.com/labstack/echo"

const (
	registerAddressPath    = "/user/addresses/register"
	getAddressByIdPath     = "/addresses/:addressId"
	getAddressesByUserPath = "/user/addresses"
	updateAddressPath      = "/addresses"
	removeAddressPath      = "/addresses"
)

type (
	AddressesRouter struct {
		Echo      *echo.Echo
		Register  echo.HandlerFunc
		GetById   echo.HandlerFunc
		GetByUser echo.HandlerFunc
		Update    echo.HandlerFunc
		Remove    echo.HandlerFunc
	}
)

func (ar *AddressesRouter) Wire() {
	ar.Echo.POST(registerAddressPath, ar.Register)
	ar.Echo.GET(getAddressByIdPath, ar.GetById)
	ar.Echo.GET(getAddressesByUserPath, ar.GetByUser)
	ar.Echo.PUT(updateAddressPath, ar.Update)
	ar.Echo.DELETE(removeAddressPath, ar.Remove)
}
