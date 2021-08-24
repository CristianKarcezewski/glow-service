package server

import "glow-service/server/postgres"

const (
	postgresDatabaseProvider = "postgres"
)

type (
	IDatabaseHandler interface {
		Select(tableName string, result interface{}, key string, value interface{}) error
		Insert(tableName string, dao interface{}) error
		Update(tableName string, dao interface{}) error
		Remove(tableName string, dao interface{}) error
		CustomQuery(tableName string, query *string) error
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
