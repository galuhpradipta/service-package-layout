package config

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func PgInit() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", os.Getenv(`PG_CONF`))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
