package routers

import "github.com/labstack/echo"

const (
	getProviderTypeByIdPath = "/providerType/:providerTypeId"
	getProviderTypesAllPath = "/providerTypesType"
)

type (
	ProviderTypesRouter struct {
		Echo    *echo.Echo
		GetById echo.HandlerFunc
		GetAll  echo.HandlerFunc
	}
)

func (pr *ProviderTypesRouter) Wire() {
	pr.Echo.GET(getProviderTypeByIdPath, pr.GetById)
	pr.Echo.GET(getProviderTypesAllPath, pr.GetAll)
}
