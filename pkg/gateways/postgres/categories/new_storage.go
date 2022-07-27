package categories

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type CategoryStorage struct {
	Conn *pgx.Pool
}

func NewStorage(connection *pgx.Pool) CategoryStorage {
	return CategoryStorage{connection}
}
