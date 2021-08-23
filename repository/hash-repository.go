package repository

import (
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryHashTable = "hashs"
)

type (
	IHashRepository interface {
		Insert(hash *dao.Hash) error
	}
	hashRepository struct {
		database server.IDatabaseHandler
	}
)

func NewHashRepository(database server.IDatabaseHandler) IHashRepository {
	return &hashRepository{database}
}

func (ur *hashRepository) Insert(hash *dao.Hash) error {
	return ur.database.Insert(repositoryHashTable, hash)
}
