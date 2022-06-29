package main

import (
	"fmt"
	"glow-service/common/functions"
	"glow-service/gateways"
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
	addressesRepository := repository.NewAddressesRepository(config.DatabaseHandler)
	userAddressesRepository := repository.NewUserAddressesRepository(config.DatabaseHandler)
	companyAddressesRepository := repository.NewCompanyAddressesRepository(config.DatabaseHandler)
	companiesRepository := repository.NewCompanyRepository(config.DatabaseHandler)
	providerTypesRepository := repository.NewProviderTypeRepository(config.DatabaseHandler)
	packagesRepository := repository.NewPackagesRepository(config.DatabaseHandler)

	// Start gateways
	locationGateway := gateways.NewLocationsGateway()
	mapsGeolocationGateway := gateways.NewMapsGeolocationGateway()

	// Start services
	authService := services.NewAuthService(config.FirebaseAuthClient)
	userService := services.NewUserService(config.FirebaseAuthClient, userRepository, hashRepository, authService)
	locationService := services.NewLocationService(locationGateway)
	mapsGeolocationService := services.NewMapsGeolocationService(mapsGeolocationGateway)
	addressesService := services.NewAddressService(addressesRepository, userAddressesRepository, companyAddressesRepository, locationService, mapsGeolocationService)
	providerTypesService := services.NewProviderTypeService(providerTypesRepository)
	packagesService := services.NewPackagesService(packagesRepository)
	companiesService := services.NewCompanyService(companiesRepository, addressesService, userService, providerTypesService, packagesService)	
	storageService := services.NewStorageService(config.FirebaseStorageClient, userService)

	// Start presenters
	authPresenter := presenters.NewAuthPresenter(config.ServerErrorMessages, authService)
	userPresenter := presenters.NewUserPresenter(config.ServerErrorMessages, authService, userService)
	locationPresenter := presenters.NewLocationPresenter(config.ServerErrorMessages, authService, locationService)
	addressesPresenter := presenters.NewAddressesPresenter(config.ServerErrorMessages, authService, addressesService, companiesService)
	companiesPresenter := presenters.NewCompanyPresenter(config.ServerErrorMessages, authService, companiesService)
	providerTypesPresenter := presenters.NewProviderTypePresenter(config.ServerErrorMessages, providerTypesService)
	packagesPresenter := presenters.NewPackagesPresenter(config.ServerErrorMessages, packagesService, authService)
	filesPresenter := presenters.NewFilePresenter(config.ServerErrorMessages, authService, storageService)

	// Start Routers
	authPresenter.Router(echo)
	userPresenter.Router(echo)
	locationPresenter.Router(echo)
	addressesPresenter.Router(echo)
	companiesPresenter.Router(echo)
	providerTypesPresenter.Router(echo)
	packagesPresenter.Router(echo)
	filesPresenter.Router(echo)
}

func initEcho(echo *echo.Echo, environment string, port int64) {
	fmt.Printf("%s => Environment %s", functions.DateToString(nil), strings.ToUpper(environment))
	_ = echo.Start(fmt.Sprintf(":%d", port))
}
