package routers

import "github.com/labstack/echo"

const (
	saveMessagesPath = "/messages"
	getMessagesPath  = "/messages"
)

type (
	MessagesRouter struct {
		Echo         *echo.Echo
		SaveMessage  echo.HandlerFunc
		FetchMessage echo.HandlerFunc
	}
)

func (mr *MessagesRouter) Wire() {
	mr.Echo.PUT(saveMessagesPath, mr.SaveMessage)
	mr.Echo.POST(getMessagesPath, mr.FetchMessage)
}
