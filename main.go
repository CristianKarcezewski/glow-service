package main

import (
	"fmt"
	"glow-service/common/functions"
	"glow-service/controllers"
	"glow-service/repository"
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
		initApplication(serverConfig, echo)
		initEcho(echo, serverConfig.Environment, int64(serverConfig.Port))
	}
}

func initApplication(config *server.Configuration, echo *echo.Echo) {
	// Start repositories
	userRepository := repository.NewUserRepository(config.DatabaseHandler)

	// Start services
	authService := services.NewAuthService(userRepository)

	// Start controllers
	authController := controllers.NewAuthController(&config.ServerErrorMessages, authService)

	// Start Routers
	authController.Router(echo, authController.Login(), authController.RefreshToken(), authController.Register()).Wire()
}

func initEcho(echo *echo.Echo, environment string, port int64) {
	fmt.Printf("%s => Environment %s", functions.DateToString(), strings.ToUpper(environment))
	echo.Start(fmt.Sprintf(":%d", port))
}
