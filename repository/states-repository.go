package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryStateTable = "states"
)

type (
	IStatesRepository interface {
		GetAll(log *models.StackLog) (*[]dao.State, error)
		GetById(log *models.StackLog, stateId int64) (*dao.State, error)
	}

	statesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewStateRepository(database server.IDatabaseHandler) IStatesRepository {
	return &statesRepository{database}
}

func (sr *statesRepository) GetAll(log *models.StackLog) (*[]dao.State, error) {
	log.AddStep("StateRepository-GetAll")

	var states []dao.State
	getErr := sr.database.GetAll(repositoryStateTable, &states)
	if getErr != nil {
		return nil, getErr
	}
	return &states, nil
}

func (sr *statesRepository) GetById(log *models.StackLog, stateId int64) (*dao.State, error) {
	log.AddStep("StateRepository-GetById")

	var state dao.State
	stateErr := sr.database.Select(repositoryStateTable, &state, "id", stateId)
	if stateErr != nil {
		return nil, stateErr
	}
	return &state, nil
}
