package routers

import "github.com/labstack/echo"

const (
	pathViacep             = "/locations/viacep/:postalCode"
	pathGetStateByUf       = "/locations/states/:stateUF"
	pathGetStates          = "/locations/states"
	pathGetCityById        = "/locations/cities/:cityId"
	pathGetCitiesByStateUF = "/locations/:stateUF/cities"
)

type (
	LocationsRouter struct {
		Echo               *echo.Echo
		ViacepAddress      echo.HandlerFunc
		GetStateByUf       echo.HandlerFunc
		GetStates          echo.HandlerFunc
		GetCityById        echo.HandlerFunc
		GetCitiesByStateUf echo.HandlerFunc
	}
)

func (lr *LocationsRouter) Wire() {
	lr.Echo.GET(pathViacep, lr.ViacepAddress)
	lr.Echo.GET(pathGetStateByUf, lr.GetStateByUf)
	lr.Echo.GET(pathGetStates, lr.GetStates)
	lr.Echo.GET(pathGetCityById, lr.GetCityById)
	lr.Echo.GET(pathGetCitiesByStateUF, lr.GetCitiesByStateUf)
}
