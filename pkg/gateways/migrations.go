package gateways

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	m, err := migrate.New(
		"file://pkg/gateways/migrations",
		"postgres://postgres:1234@localhost:5432/challenge?sslmode=disable")
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
