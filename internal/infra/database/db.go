package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moura95/go-ddd/internal/infra/cfg"
)

type Connection interface {
	Close() error
	DB() *sqlx.DB
}

type conn struct {
	db *sqlx.DB
}

func ConnectPostgres() (Connection, error) {
	loadConfig, _ := cfg.LoadConfig(".")

	db, err := sqlx.Open("postgres", loadConfig.DBSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &conn{db: db}, nil
}

func (c *conn) Close() error {
	return c.db.Close()
}

func (c *conn) DB() *sqlx.DB {
	return c.db
}
