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
	stateService struct {
		stateRepository repository.IStateRepository
	}
)

func NewStateService(stateRepository repository.IStateRepository) IStatesService {
	return &stateService{stateRepository}
}

func (ss *stateService) GetAll(log *models.StackLog) (*[]models.State, error) {
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

func (ss *stateService) GetById(log *models.StackLog, stateId int64) (*models.State, error) {
	log.AddStep("StateService-GetById")

	result, resultErr := ss.stateRepository.GetById(log, stateId)
	if resultErr != nil {
		return nil, resultErr
	}
	return result.ToModel(), nil
}
