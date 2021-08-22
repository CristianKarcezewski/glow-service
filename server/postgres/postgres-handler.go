package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type (
	postgresHandler struct {
		con *pg.DB
	}
)

func NewPostgresHandler(databaseProvider, user, password, address, port, database *string) *postgresHandler {
	addr := fmt.Sprintf("%s:%s", *address, *port)
	options := &pg.Options{
		User:     *user,
		Password: *password,
		Addr:     addr,
		Database: *databaseProvider,
		PoolSize: 50,
	}

	con := pg.Connect(options)
	if con == nil {
		panic("unable to connect to database")
	}

	pingErr := con.Ping(context.Background())
	if pingErr != nil {
		panic("database connection test error")
	}
	return &postgresHandler{con}
}

func (p *postgresHandler) Insert(tableName string, dao interface{}) error {
	obj, err := p.con.Model(dao).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
	return err
}

func (p *postgresHandler) FindById(tableName string, userId *int64, dao interface{}) (interface{}, error) {
	err := p.con.Model(dao).Where(fmt.Sprintf("id = %d", *userId)).Select()
	return dao, err
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

func (p *postgresHandler) CustomQuery(tableName string, query *string) error {
	return nil
}
