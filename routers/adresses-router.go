package routers

import "github.com/labstack/echo"

const (
	registerAddressUserPath    = "/user/address/register"
	registerAddressCompanyPath = "/company/address/register"
	getAddressByIdPath         = "/address/:addressId"
	getAddressesByUserPath     = "/user/addresses"
	getAddressesByCompanyPath  = "/company/addresses"
	updateAddressPath          = "/address"
	removeAddressPath          = "/address/:addressId"
)

type (
	AddressesRouter struct {
		Echo         *echo.Echo
		Register     echo.HandlerFunc
		GetById      echo.HandlerFunc
		GetByUser    echo.HandlerFunc
		GetByCompany echo.HandlerFunc
		Update       echo.HandlerFunc
		Remove       echo.HandlerFunc
	}
)

func (ar *AddressesRouter) Wire() {
	ar.Echo.POST(registerAddressUserPath, ar.Register)
	ar.Echo.POST(registerAddressCompanyPath, ar.Register)
	ar.Echo.GET(getAddressByIdPath, ar.GetById)
	ar.Echo.GET(getAddressesByUserPath, ar.GetByUser)
	ar.Echo.PUT(updateAddressPath, ar.Update)
	ar.Echo.DELETE(removeAddressPath, ar.Remove)
}
