package postgres

import (
	"context"
	"errors"
	"fmt"
	"glow-service/common/functions"

	"github.com/jackc/pgx/v4"
)

type (
	postgresHandler struct {
		conn *pgx.Conn
	}
)

func NewPostgresHandler(databaseProvider, user, password, address, port, database *string) *postgresHandler {
	conn, _ := pgx.Connect(context.Background(), fmt.Sprintf("%s://%s:%s@%s:%s/%s", *databaseProvider, *user, *password, *address, *port, *database))

	connErr := conn.Ping(context.Background())
	if connErr != nil {
		conn.Close(context.Background())
		panic(connErr)
	}

	conn.Close(context.Background())
	return &postgresHandler{conn}
}

func (p *postgresHandler) Insert(tableName string, dao interface{}) error {
	query := fmt.Sprintf("INSERT INTO %s", tableName)
	keysError := p.mountQuery(&query, dao)
	if keysError != nil {
		return keysError
	}
	if _, err := p.conn.Exec(context.Background(), query, dao); err != nil {
		return err
	}
	return nil
}

func (p *postgresHandler) Select(tableName string, dao interface{}) error {
	return nil
}

func (p *postgresHandler) Update(tableName string, dao interface{}) error {
	return nil
}

func (p *postgresHandler) Remove(tableName string, dao interface{}) error {
	return nil
}

func (p *postgresHandler) CustomQuery(string) error {
	return nil
}

func (p *postgresHandler) mountQuery(query *string, dao interface{}) error {
	var strKeys []string
	m, ok := dao.(map[string]interface{})
	if !ok {
		return errors.New("error mapping DAO object")
	}
	for key, _ := range m {
		if len(strKeys) == 0 {
			*query = fmt.Sprintf("%s(%s", *query, functions.ToSnakeCase(key))
		} else {
			*query = fmt.Sprintf("%s, %s", *query, functions.ToSnakeCase(key))
		}
		strKeys = append(strKeys, key)
	}

	*query = fmt.Sprintf("%s) VALUES(", *query)
	for index := range strKeys {
		if index == 0 {
			*query = fmt.Sprint("$%d", (index + 1))
		} else {
			*query = fmt.Sprint(", $%d", (index + 1))
		}
	}
	*query = fmt.Sprintf("%s)", *query)
	return nil
}
