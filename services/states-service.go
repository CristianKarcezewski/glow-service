package services

import (
	"glow-service/models"
	"glow-service/repository"
)

type (
	IStatesService interface {
		GetAll(log *models.StackLog) (*[]models.State, error)
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
