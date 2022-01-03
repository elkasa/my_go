package repository

import (
	"github.com/jackc/pgx/v4"
)

type idfRepositoryPgx struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *idfRepositoryPgx {
	return &idfRepositoryPgx{db}
}
