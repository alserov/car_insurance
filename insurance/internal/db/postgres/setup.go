package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func MustConnect(addr, migrations string) *sqlx.DB {
	conn, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	mustMigrate(conn, migrations)

	return conn
}

func mustMigrate(conn *sqlx.DB, path string) {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(conn.DB, path); err != nil {
		panic(err)
	}
}
