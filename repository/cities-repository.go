package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryCitiesTable = "cities"
)

type (
	ICitiesRepository interface {
		GetById(log *models.StackLog, cityId int64) (*dao.City, error)
		GetByState(log *models.StackLog, stateId int64) (*[]dao.City, error)
	}
	citiesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewCitiesRepository(database server.IDatabaseHandler) ICitiesRepository {
	return &citiesRepository{database}
}

func (cr *citiesRepository) GetById(log *models.StackLog, cityId int64) (*dao.City, error) {
	log.AddStep("CitiesRepository-GetByState")
	var city dao.City
	cityError := cr.database.Select(repositoryCitiesTable, &city, "id", cityId)
	if cityError != nil {
		return nil, cityError
	}
	return &city, nil
}

func (cr *citiesRepository) GetByState(log *models.StackLog, stateId int64) (*[]dao.City, error) {
	log.AddStep("CitiesRepository-GetByState")
	var cities []dao.City
	citiesError := cr.database.Select(repositoryCitiesTable, &cities, "state_id", stateId)
	if citiesError != nil {
		return nil, citiesError
	}
	return &cities, nil
}
