package main

import (
	"fmt"
	"glow-service/common/functions"
	"glow-service/controllers"
	"glow-service/models"
	"glow-service/server"
	"glow-service/services"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	serverConfig, serverConfigError := server.ConfigurationInstance()
	if serverConfigError != nil {
		fmt.Println(serverConfigError.Error())
	} else {
		echo := echo.New()
		initApplication(&serverConfig.ServerErrorMessages, echo)
		initEcho(echo, serverConfig.Environment, int64(serverConfig.Port))
	}
}

func initApplication(errorMessages *models.ServerErrorMessages, echo *echo.Echo) {
	// Start services
	authService := services.NewAuthService()

	// Start controllers
	authController := controllers.NewAuthController(errorMessages, authService)

	// Start Routers
	authController.Router(echo, authController.Login(), authController.RefreshToken()).Wire()
}

func initEcho(echo *echo.Echo, environment string, port int64) {
	fmt.Printf("%s => Environment %s", functions.DateToString(), strings.ToUpper(environment))
	echo.Start(fmt.Sprintf(":%d", port))
}
