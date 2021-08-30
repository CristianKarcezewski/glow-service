package services

import (
	"glow-service/models"
	"glow-service/repository"
)

type (
	ICitiesService interface {
		GetById(log *models.StackLog, cityId int64) (*models.City, error)
		GetByState(log *models.StackLog, stateId int64) (*[]models.City, error)
	}
	citiesService struct {
		citiesRepository repository.ICitiesRepository
	}
)

func NewCitiesService(citiesRepository repository.ICitiesRepository) ICitiesService {
	return &citiesService{citiesRepository}
}

func (cs *citiesService) GetById(log *models.StackLog, cityId int64) (*models.City, error) {
	log.AddStep("CitiesService-GetById")

	city, cityErr := cs.citiesRepository.GetById(log, cityId)
	if cityErr != nil {
		return nil, cityErr
	}
	return city.ToModel(), nil
}

func (cs *citiesService) GetByState(log *models.StackLog, stateId int64) (*[]models.City, error) {
	log.AddStep("CitiesService-GetByState")

	var cities []models.City
	citiesRepository, citiesError := cs.citiesRepository.GetByState(log, stateId)
	if citiesError != nil {
		return nil, citiesError
	}

	for i := range *citiesRepository {
		cities = append(cities, *(*citiesRepository)[i].ToModel())
	}

	return &cities, nil
}
