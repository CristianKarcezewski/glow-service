package routers

import "github.com/labstack/echo"

const (
	getPackagesByIdPath = "/packages/:packageId"
	getAllPackagesPath  = "/packages"
)

type (
	PackagesRouter struct {
		Echo    *echo.Echo
		GetById echo.HandlerFunc
		GetAll  echo.HandlerFunc
	}
)

func (pr *PackagesRouter) Wire() {
	pr.Echo.GET(getPackagesByIdPath, pr.GetById)
	pr.Echo.GET(getAllPackagesPath, pr.GetAll)
}
