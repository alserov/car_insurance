package postgres

import "github.com/jmoiron/sqlx"

func MustConnect(addr string) *sqlx.DB {
	conn, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	return conn
}
