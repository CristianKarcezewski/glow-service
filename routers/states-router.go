package routers

import "github.com/labstack/echo"

const (
	statesPath = "/states"
)

type (
	StatesRouter struct {
		Echo   *echo.Echo
		GetAll echo.HandlerFunc
	}
)

func (sr *StatesRouter) Wire() {
	sr.Echo.GET(statesPath, sr.GetAll)
}
