package gateways

import (
	"encoding/json"
	"fmt"
	"glow-service/models"
	"glow-service/models/response"
	"net/http"
	"strings"
)

const (
	mapsKey         = "AIzaSyCIP8t2P7CWthvCdhCsLJrB6KfZ-lgnY14"
	mapsGeolocation = "https://maps.googleapis.com/maps/api/geocode/json"
)

type (
	IMapsGeolocationGateway interface {
		FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*response.MapsResponse, error)
		FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*response.MapsResponse, error)
	}

	mapsGeolocationGateway struct{}
)

func NewMapsGeolocationGateway() IMapsGeolocationGateway {
	return &mapsGeolocationGateway{}
}

func (mg *mapsGeolocationGateway) FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*response.MapsResponse, error) {
	log.AddStep("MapsGeolocationGateway-FindAddressByGeolocation")

	uri := fmt.Sprintf("%s?latlng=%s,%s&key=%s", mapsGeolocation, geolocation.Latitude, geolocation.Longitude, mapsKey)
	log.AddInfo(fmt.Sprintf("MapsGeolocationGateway - Requesting maps address info (%s, %s)", geolocation.Latitude, geolocation.Longitude))
	resp, respErr := http.Get(uri)
	if respErr != nil {
		return nil, respErr
	}

	var mapsResponse response.MapsResponse
	decodeErr := json.NewDecoder(resp.Body).Decode(&mapsResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &mapsResponse, nil
}

func (mg *mapsGeolocationGateway) FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*response.MapsResponse, error) {
	log.AddStep("MapsGeolocationGateway-FindGeolocationByAddress")

	uri := fmt.Sprintf("%s?address=%s,&key=%s", mapsGeolocation, strings.ReplaceAll(geolocation.Street, " ", "%20"), mapsKey)
	log.AddInfo(fmt.Sprintf("MapsGeolocationGateway - Requesting maps address info (%s, %s)", geolocation.Latitude, geolocation.Longitude))
	resp, respErr := http.Get(uri)
	if respErr != nil {
		return nil, respErr
	}

	var mapsResponse response.MapsResponse
	decodeErr := json.NewDecoder(resp.Body).Decode(&mapsResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	resp.Body.Close()
	return &mapsResponse, nil
}
