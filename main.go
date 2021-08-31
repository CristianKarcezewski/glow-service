package main

import (
	"fmt"
	"glow-service/common/functions"
	"glow-service/controllers"
	"glow-service/repository"
	"glow-service/server"
	"glow-service/services"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	serverConfig, serverConfigError := server.ConfigurationInstance()
	if serverConfigError != nil {
		fmt.Println(serverConfigError.Error())
	} else {
		e := echo.New()
		initApplication(serverConfig, e)
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "You found me!")
		})
		initEcho(e, serverConfig.Environment, int64(serverConfig.Port))
	}
}

func initApplication(config *server.Configuration, echo *echo.Echo) {
	// Start repositories
	userRepository := repository.NewUserRepository(config.DatabaseHandler)
	hashRepository := repository.NewHashRepository(config.DatabaseHandler)
	statesRepository := repository.NewStateRepository(config.DatabaseHandler)
	citiesRepository := repository.NewCitiesRepository(config.DatabaseHandler)
	addressesRepository := repository.NewAddressesRepository(config.DatabaseHandler)
	userAddressesRepository := repository.NewUserAddressesRepository(config.DatabaseHandler)

	// Start services
	userService := services.NewUserService(userRepository, hashRepository)
	authService := services.NewAuthService(userService)
	statesService := services.NewStateService(statesRepository)
	citiesService := services.NewCitiesService(citiesRepository)
	addressesService := services.NewAddressService(addressesRepository, userAddressesRepository, statesService, citiesService)

	// Start controllers
	authController := controllers.NewAuthController(&config.ServerErrorMessages, authService)
	statesController := controllers.NewStatesController(&config.ServerErrorMessages, statesService)
	citiesController := controllers.NewCitiesController(&config.ServerErrorMessages, citiesService)
	addressesController := controllers.NewAddressesController(&config.ServerErrorMessages, authService, addressesService)

	// Start Routers
	authController.Router(echo, authController.Login(), authController.RefreshToken(), authController.Register()).Wire()
	statesController.Router(echo, statesController.GetAll()).Wire()
	citiesController.Router(echo, citiesController.GetById(), citiesController.GetByState()).Wire()
	addressesController.Router(echo, addressesController.Register()).Wire()
}

func initEcho(echo *echo.Echo, environment string, port int64) {
	fmt.Printf("%s => Environment %s", functions.DateToString(), strings.ToUpper(environment))
	echo.Start(fmt.Sprintf(":%d", port))
}
