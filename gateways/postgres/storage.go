package postgres

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type VideoStorage struct {
	Conn *pgx.Pool
}

func NewStorage(connection *pgx.Pool) VideoStorage {
	return VideoStorage{connection}
}
