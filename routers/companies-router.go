package routers

import "github.com/labstack/echo"

const (
	registerCompanyPath = "/companies"
	getCompanyByIdPath  = "/companies"
	getByUserPath       = "/companies/user"
	updateCompanyPath   = "/companies"
	searchCompanyPath   = "/companies/search"
)

type (
	CompaniesRouter struct {
		Echo      *echo.Echo
		Register  echo.HandlerFunc
		GetById   echo.HandlerFunc
		GetByUser echo.HandlerFunc
		Update    echo.HandlerFunc
		// Remove   echo.HandlerFunc
		Search echo.HandlerFunc
	}
)

func (cr *CompaniesRouter) Wire() {
	cr.Echo.POST(registerCompanyPath, cr.Register)
	cr.Echo.GET(getCompanyByIdPath, cr.GetById)
	cr.Echo.GET(getByUserPath, cr.GetByUser)
	cr.Echo.PUT(updateCompanyPath, cr.Update)
	// cr.Echo.DELETE(removeCompanyPath, cr.Remove)
	cr.Echo.POST(searchCompanyPath, cr.Search)
}
