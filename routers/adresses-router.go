package routers

import "github.com/labstack/echo"

const (
	pathGetById = "addresses/:addressId"
	pathUpdate  = "addresses"

	pathGetByUser      = "addresses/user"
	pathRegisterByUser = "addresses/user"
	pathRemoveByUser   = "addresses/user/:addressId"

	pathGetByCompany      = "addresses/company"
	pathRegisterByCompany = "addresses/company"
	pathRemoveByCompany   = "addresses/company/:addressId"
)

type (
	AddressesRouter struct {
		Echo              *echo.Echo
		GetById           echo.HandlerFunc
		GetByUser         echo.HandlerFunc
		GetByCompany      echo.HandlerFunc
		RegisterByUser    echo.HandlerFunc
		RegisterByCompany echo.HandlerFunc
		Update            echo.HandlerFunc
		RemoveByUser      echo.HandlerFunc
		RemoveByCompany   echo.HandlerFunc
	}
)

func (ar *AddressesRouter) Wire() {
	ar.Echo.GET(pathGetById, ar.GetById)
	ar.Echo.GET(pathGetByUser, ar.GetByUser)
	ar.Echo.GET(pathGetByCompany, ar.GetByCompany)
	ar.Echo.POST(pathRegisterByUser, ar.RegisterByUser)
	ar.Echo.POST(pathRegisterByCompany, ar.RegisterByCompany)
	ar.Echo.PUT(pathUpdate, ar.Update)
	ar.Echo.DELETE(pathRemoveByUser, ar.RemoveByUser)
	ar.Echo.DELETE(pathRemoveByCompany, ar.RemoveByCompany)
}
