package main

import (
	"fmt"
	"glow-service/common/functions"
	"glow-service/presenters"
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
	authPresenter := presenters.NewAuthPresenter(&config.ServerErrorMessages, authService)
	statesPresenter := presenters.NewStatesPresenter(&config.ServerErrorMessages, statesService)
	citiesPresenter := presenters.NewCitiesPresenter(&config.ServerErrorMessages, citiesService)
	addressesPresenter := presenters.NewAddressesPresenter(&config.ServerErrorMessages, authService, addressesService)

	// Start Routers
	authPresenter.Router(echo, authPresenter.Login(), authPresenter.RefreshToken(), authPresenter.Register()).Wire()
	statesPresenter.Router(echo, statesPresenter.GetAll()).Wire()
	citiesPresenter.Router(echo, citiesPresenter.GetById(), citiesPresenter.GetByState()).Wire()
	addressesPresenter.Router(echo, addressesPresenter.Register(), addressesPresenter.GetById(), addressesPresenter.GetByUser(), addressesPresenter.Update(), addressesPresenter.Remove()).Wire()
}

func initEcho(echo *echo.Echo, environment string, port int64) {
	fmt.Printf("%s => Environment %s", functions.DateToString(), strings.ToUpper(environment))
	_ = echo.Start(fmt.Sprintf(":%d", port))
}
