package services

import (
	"glow-service/models"
	"glow-service/repository"
)

type (
	IStatesService interface {
		GetAll(log *models.StackLog) (*[]models.State, error)
		GetById(log *models.StackLog, stateId int64) (*models.State, error)
	}
	statesService struct {
		stateRepository repository.IStatesRepository
	}
)

func NewStateService(stateRepository repository.IStatesRepository) IStatesService {
	return &statesService{stateRepository}
}

func (ss *statesService) GetAll(log *models.StackLog) (*[]models.State, error) {
	log.AddStep("StateService-GetAll")

	var states []models.State

	repositoryStates, err := ss.stateRepository.GetAll(log)
	if err != nil {
		return nil, err
	}

	for i := range *repositoryStates {
		states = append(states, *(*repositoryStates)[i].ToModel())
	}

	return &states, nil
}

func (ss *statesService) GetById(log *models.StackLog, stateId int64) (*models.State, error) {
	log.AddStep("StateService-GetById")

	result, resultErr := ss.stateRepository.GetById(log, stateId)
	if resultErr != nil {
		return nil, resultErr
	}
	return result.ToModel(), nil
}
