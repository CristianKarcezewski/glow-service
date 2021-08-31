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
		Database: *database,
	}

	con := pg.Connect(options)
	if con == nil {
		panic("unable to connect to database")
	}

	pingErr := con.Ping(context.Background())
	if pingErr != nil {
		fmt.Print(pingErr)
		panic("database connection test error")
	}
	return &postgresHandler{con}
}

func (p *postgresHandler) GetAll(tableName string, result interface{}) error {
	return p.con.Model(result).Select()
}

func (p *postgresHandler) Insert(tableName string, dao interface{}) error {
	// p.createTable(dao)
	_, err := p.con.Model(dao).Returning("*").Insert(dao)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresHandler) Select(tableName string, result interface{}, key string, value interface{}) error {
	// var key string
	// var value interface{}

	//dynamically extracts the filter value
	// var mobj map[string]interface{}
	// inrec, _ := json.Marshal(filter)
	// json.Unmarshal(inrec, &mobj)
	// for k, v := range mobj {
	// 	key = k
	// 	value = v
	// 	break
	// }

	return p.con.Model(result).Where(fmt.Sprintf("%s = ?", key), value).Select()
	// return p.con.Model(filter).Limit(1).Select(&mobj)
}

func (p *postgresHandler) Update(tableName string, dao interface{}) error {
	_, err := p.con.Model(dao).WherePK().Returning("*").Update(dao)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresHandler) Remove(tableName string, dao interface{}, key string, value interface{}) error {
	_, err := p.con.Model(dao).Where(fmt.Sprintf("%s = ?", key), value).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresHandler) CustomQuery() (*pg.DB, error) {
	return p.con, nil
}

// func (p *postgresHandler) createTable(dao interface{}) {
// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}

// 	createErr := p.con.Model(dao).CreateTable(opts)
// 	if createErr != nil {
// 		panic(createErr)
// 	}
// }
