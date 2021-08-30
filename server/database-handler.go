package server

import (
	"glow-service/server/postgres"

	"github.com/go-pg/pg/v10"
)

const (
	postgresDatabaseProvider = "postgres"
)

type (
	IDatabaseHandler interface {
		GetAll(tableName string, result interface{}) error
		Select(tableName string, result interface{}, key string, value interface{}) error
		Insert(tableName string, dao interface{}) error
		Update(tableName string, dao interface{}) error
		Remove(tableName string, dao interface{}, key string, value interface{}) error
		CustomQuery() (*pg.DB, error)
	}
)

func SetubDatabase(databaseProvider, user, password, address, port, database *string) IDatabaseHandler {
	switch *databaseProvider {
	case postgresDatabaseProvider:
		return postgres.NewPostgresHandler(databaseProvider, user, password, address, port, database)
		break
	}
	return nil
}
