package routers

import "github.com/labstack/echo"

const (
	getCityByIdPath      = "/cities/:cityId"
	getCitiesByStatePath = "/cities"
)

type (
	CitiesRouter struct {
		Echo       *echo.Echo
		GetById    echo.HandlerFunc
		GetByState echo.HandlerFunc
	}
)

func (cr *CitiesRouter) Wire() {
	cr.Echo.GET(getCityByIdPath, cr.GetById)
	cr.Echo.GET(getCitiesByStatePath, cr.GetByState)
}
