package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryHashTable = "hashs"
)

type (
	IHashRepository interface {
		Insert(log *models.StackLog, hash *dao.Hash) error
		Select(log *models.StackLog, key string, value interface{}) (*dao.Hash, error)
	}
	hashRepository struct {
		database server.IDatabaseHandler
	}
)

func NewHashRepository(database server.IDatabaseHandler) IHashRepository {
	return &hashRepository{database}
}

func (hr *hashRepository) Insert(log *models.StackLog, hash *dao.Hash) error {
	log.AddStep("HashRepository-Insert")

	log.AddInfo("Saving encrypted password")
	return hr.database.Insert(repositoryHashTable, hash)
}

func (hr *hashRepository) Select(log *models.StackLog, key string, value interface{}) (*dao.Hash, error) {
	log.AddStep("HashRepository-Select")

	var hash dao.Hash

	log.AddInfo("validating password")
	selectErr := hr.database.Select(repositoryHashTable, &hash, key, value)
	if selectErr != nil {
		return nil, selectErr
	}
	return &hash, nil
}
