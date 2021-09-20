package gateways

import (
	"encoding/json"
	"fmt"
	"glow-service/models"
	"glow-service/models/response"
	"net/http"
)

const (
	ibgeStates = "https://servicodados.ibge.gov.br/api/v1/localidades/estados"
	ibgeCities = "https://servicodados.ibge.gov.br/api/v1/localidades/municipios"
	viacep     = "https://viacep.com.br/ws"
)

type (
	ILocationsGateway interface {
		GetViacep(log *models.StackLog, postalCode string) (*response.Viacep, error)
		GetStateByUF(log *models.StackLog, stateUF string) (*response.State, error)
		GetStates(log *models.StackLog) (*[]response.State, error)
		GetCityById(log *models.StackLog, cityId int64) (*response.City, error)
		GetCitiesByState(log *models.StackLog, stateUF string) (*[]response.City, error)
	}
	locationsGateway struct{}
)

func NewLocationsGateway() ILocationsGateway {
	return &locationsGateway{}
}

func (lg *locationsGateway) GetViacep(log *models.StackLog, postalCode string) (*response.Viacep, error) {
	log.AddStep("LocationsGateway-GetViacep")

	resp, respErr := http.Get(fmt.Sprintf("%s/%s/json", viacep, postalCode))
	if respErr != nil {
		return nil, respErr
	}

	var viacep response.Viacep
	decodeErr := json.NewDecoder(resp.Body).Decode(&viacep)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &viacep, nil
}

func (lg *locationsGateway) GetStateByUF(log *models.StackLog, stateUF string) (*response.State, error) {
	log.AddStep("LocationsGateway-GetStateById")

	resp, respErr := http.Get(fmt.Sprintf("%s/%s", ibgeStates, stateUF))
	if respErr != nil {
		return nil, respErr
	}

	var state response.State
	decodeErr := json.NewDecoder(resp.Body).Decode(&state)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &state, nil
}

func (lg *locationsGateway) GetStates(log *models.StackLog) (*[]response.State, error) {
	log.AddStep("LocationsGateway-GetStates")

	resp, respErr := http.Get(ibgeStates)
	if respErr != nil {
		return nil, respErr
	}

	var states []response.State
	decodeErr := json.NewDecoder(resp.Body).Decode(&states)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &states, nil
}

func (lg *locationsGateway) GetCityById(log *models.StackLog, cityId int64) (*response.City, error) {
	log.AddStep("LocationsGateway-GetCityById")

	resp, respErr := http.Get(fmt.Sprintf("%s/%d", ibgeCities, cityId))
	if respErr != nil {
		return nil, respErr
	}

	var city response.City
	decodeErr := json.NewDecoder(resp.Body).Decode(&city)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &city, nil
}

func (lg *locationsGateway) GetCitiesByState(log *models.StackLog, stateUF string) (*[]response.City, error) {
	log.AddStep("LocationsGateway-GetCityById")

	resp, respErr := http.Get(fmt.Sprintf("%s/%s/municipios", ibgeStates, stateUF))
	if respErr != nil {
		return nil, respErr
	}

	var cities []response.City
	decodeErr := json.NewDecoder(resp.Body).Decode(&cities)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &cities, nil
}
