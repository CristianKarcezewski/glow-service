package services

import (
	"fmt"
	"glow-service/gateways"
	"glow-service/models"
	"glow-service/models/response"
	"strconv"
	"sync"
)

type (
	IMapsGeolocationService interface {
		FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*models.Address, error)
		FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*models.Address, error)
	}

	mapsGeolocationService struct {
		MapsGeolocationGateway gateways.IMapsGeolocationGateway
	}
)

func NewMapsGeolocationService(mapsGeolocationGateway gateways.IMapsGeolocationGateway) IMapsGeolocationService {
	return &mapsGeolocationService{mapsGeolocationGateway}
}

func (mg *mapsGeolocationService) FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*models.Address, error) {
	log.AddStep("MapsGeolocationService-FindAddressByGeolocation")

	mapsResponse, mapsResponseError := mg.MapsGeolocationGateway.FindAddressByGeolocation(log, geolocation)
	if mapsResponseError != nil {
		return nil, mapsResponseError
	}

	wg := sync.WaitGroup{}
	wg.Add(6)
	mg.fetchState(&wg, log, mapsResponse, geolocation)
	mg.fetchCity(&wg, log, mapsResponse, geolocation)
	mg.fetchDistrict(&wg, log, mapsResponse, geolocation)
	mg.fetchStreet(&wg, log, mapsResponse, geolocation)
	mg.fetchNumber(&wg, log, mapsResponse, geolocation)
	mg.fetchPostalCode(&wg, log, mapsResponse, geolocation)
	wg.Wait()

	return geolocation, nil
}

func (mg *mapsGeolocationService) FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*models.Address, error) {
	log.AddStep("MapsGeolocationService-FindGeolocationByAddress")

	mapsResponse, mapsResponseError := mg.MapsGeolocationGateway.FindGeolocationByAddress(log, geolocation)
	if mapsResponseError != nil {
		return nil, mapsResponseError
	}

	wg := sync.WaitGroup{}
	wg.Add(6)
	mg.fetchPostalCode(&wg, log, mapsResponse, geolocation)
	wg.Wait()

	for _, addressComponent := range mapsResponse.Results {
		geolocation.Latitude = fmt.Sprintf("%f", addressComponent.Geometry.Location.Lat)
		geolocation.Longitude = fmt.Sprintf("%f", addressComponent.Geometry.Location.Lng)
	}

	return geolocation, nil
}

func (mg *mapsGeolocationService) fetchState(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	secondCondition := false
	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false
			secondCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "administrative_area_level_1" {
					firstCondition = true
				}
				if cType == "political" {
					secondCondition = true
				}
			}

			if firstCondition && secondCondition {
				address.State.Name = responseComponent.LongName
				address.State.Uf = responseComponent.ShortName
				wg.Done()
				return
			}
		}
	}

	wg.Done()
}

func (mg *mapsGeolocationService) fetchCity(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	secondCondition := false
	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false
			secondCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "administrative_area_level_2" {
					firstCondition = true
				}
				if cType == "political" {
					secondCondition = true
				}
			}

			if firstCondition && secondCondition {
				address.City.Name = responseComponent.LongName
			}
		}
	}

	wg.Done()
}

func (mg *mapsGeolocationService) fetchDistrict(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	secondCondition := false
	thirdCondition := false

	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false
			secondCondition = false
			thirdCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "political" {
					firstCondition = true
				}
				if cType == "sublocality" {
					secondCondition = true
				}
				if cType == "sublocality_level_1" {
					thirdCondition = true
				}
			}

			if firstCondition && secondCondition && thirdCondition {
				address.District = responseComponent.LongName
				wg.Done()
				return
			}
		}
	}

	wg.Done()
}

func (mg *mapsGeolocationService) fetchStreet(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "route" {
					firstCondition = true
				}
			}

			if firstCondition {
				address.Street = responseComponent.LongName
				wg.Done()
				return
			}
		}
	}

	wg.Done()
}

func (mg *mapsGeolocationService) fetchNumber(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "street_number" {
					firstCondition = true
				}
			}

			if firstCondition {
				i, err := strconv.ParseInt(responseComponent.LongName, 10, 64)
				if err == nil {
					address.Number = i
					wg.Done()
					return
				}
			}
		}
	}

	wg.Done()
}

func (mg *mapsGeolocationService) fetchPostalCode(wg *sync.WaitGroup, log *models.StackLog, mapsResponse *response.MapsResponse, address *models.Address) {
	firstCondition := false
	for _, responseComponents := range mapsResponse.Results {
		for _, responseComponent := range responseComponents.Components {
			firstCondition = false

			for _, cType := range responseComponent.Types {
				if cType == "postal_code" {
					firstCondition = true
				}
			}

			if firstCondition {
				address.PostalCode = responseComponent.LongName
				wg.Done()
				return
			}
		}
	}

	wg.Done()
}
