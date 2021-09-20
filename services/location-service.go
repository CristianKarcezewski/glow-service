package services

import (
	"errors"
	"glow-service/gateways"
	"glow-service/models"
	"strconv"
	"sync"
)

type (
	ILocationService interface {
		FindByPostalCode(log *models.StackLog, postalCode string) (*models.Address, error)
		FindStateByUf(log *models.StackLog, stateUF string) (*models.State, error)
		FindStates(log *models.StackLog) (*[]models.State, error)
		FindCityById(log *models.StackLog, cityId int64) (*models.City, error)
		FindCitiesByStateUf(log *models.StackLog, stateUF string) (*[]models.City, error)
	}

	locationService struct {
		locationGateway gateways.ILocationsGateway
	}
)

func NewLocationService(locationGateway gateways.ILocationsGateway) ILocationService {
	return &locationService{locationGateway}
}

func (ls *locationService) FindByPostalCode(log *models.StackLog, postalCode string) (*models.Address, error) {
	log.AddStep("LocationService-FindByViacep")

	viacep, viacepErr := ls.locationGateway.GetViacep(log, postalCode)
	if viacepErr != nil {
		return nil, viacepErr
	}
	ctId, ctErr := strconv.ParseInt(viacep.Ibge, 10, 64)
	if ctErr != nil {
		return nil, ctErr
	}

	var state models.State
	var city models.City
	var err string
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go ls.findStateByUfAsync(wg, log, viacep.Uf, &state, &err)
	go ls.findCityByIdAsync(wg, log, ctId, &city, &err)
	wg.Wait()

	if err != "" {
		return nil, errors.New(err)
	}

	return &models.Address{
		PostalCode: viacep.Cep,
		StateUF:    state.Uf,
		CityId:     city.CityId,
		District:   viacep.District,
		Street:     viacep.Street,
		State:      state,
		City:       city,
	}, nil
}

func (ls *locationService) FindStateByUf(log *models.StackLog, stateUF string) (*models.State, error) {
	log.AddStep("LocationService-FindStateByUf")

	state, err := ls.locationGateway.GetStateByUF(log, stateUF)
	if err != nil {
		return nil, err
	}
	return state.ToModel(), nil
}

func (ls *locationService) FindStates(log *models.StackLog) (*[]models.State, error) {
	log.AddStep("LocationService-FindStates")

	statesResponse, respErr := ls.locationGateway.GetStates(log)
	if respErr != nil {
		return nil, respErr
	}

	var states []models.State
	for i := range *statesResponse {
		states = append(states, *(*statesResponse)[i].ToModel())
	}

	return &states, nil
}

func (ls *locationService) FindCityById(log *models.StackLog, cityId int64) (*models.City, error) {
	log.AddStep("LocationService-FindCityById")

	city, err := ls.locationGateway.GetCityById(log, cityId)
	if err != nil {
		return nil, err
	}

	return city.ToModel(), nil
}

func (ls *locationService) FindCitiesByStateUf(log *models.StackLog, stateUF string) (*[]models.City, error) {
	log.AddStep("LocationService-FindCityById")

	citiesResponse, respErr := ls.locationGateway.GetCitiesByState(log, stateUF)
	if respErr != nil {
		return nil, respErr
	}

	var cities []models.City
	for i := range *citiesResponse {
		cities = append(cities, *(*citiesResponse)[i].ToModel())
	}

	return &cities, nil
}

func (ls *locationService) findStateByUfAsync(wg *sync.WaitGroup, log *models.StackLog, stateUF string, state *models.State, err *string) {
	stRes, stErr := ls.locationGateway.GetStateByUF(log, stateUF)
	if stErr != nil {
		*err = stErr.Error()
	} else {
		state.StateId = stRes.Id
		state.Uf = stRes.Uf
		state.Name = stRes.Name
	}
	wg.Done()
}

func (ls *locationService) findCityByIdAsync(wg *sync.WaitGroup, log *models.StackLog, cityId int64, city *models.City, err *string) {
	ctRes, ctErr := ls.locationGateway.GetCityById(log, cityId)
	if ctErr != nil {
		*err = ctErr.Error()
	} else {
		city.CityId = ctRes.Id
		city.Name = ctRes.Name
	}
	wg.Done()
}
