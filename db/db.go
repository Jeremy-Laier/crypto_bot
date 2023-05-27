package db

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"os"
)

type dbImpl struct {
	*sql.DB
}

func New() DB {
	connstring := os.Getenv("PG_CONN_STRING")
	underlyingDB, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal().Err(err)
	}

	return &dbImpl{underlyingDB}
}
