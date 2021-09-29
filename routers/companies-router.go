package routers

import "github.com/labstack/echo"

const (
	registerCompanyPath = "/companies"
)

type (
	CompaniesRouter struct {
		Echo     *echo.Echo
		Register echo.HandlerFunc
		// GetById  echo.HandlerFunc
		// Update   echo.HandlerFunc
		// Remove   echo.HandlerFunc
	}
)

func (cr *CompaniesRouter) Wire() {
	cr.Echo.POST(registerCompanyPath, cr.Register)
	// cr.Echo.GET(getCompanyByIdPath, cr.GetById)
	// cr.Echo.PUT(updateCompanyPath, cr.Update)
	// cr.Echo.DELETE(removeCompanyPath, cr.Remove)
}
