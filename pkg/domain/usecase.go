package domain

import (
	"context"
	"errors"

	"github.com/daniel1sender/alura-flix/pkg/domain/entity"
)

var (
	ErrNoVideoFound = errors.New("no video found")
)

type UseCase interface {
	GetAll(ctx context.Context) ([]entity.Video, error)
	GetByID(ctx context.Context, id string) (entity.Video, error)
	Create(ctx context.Context, title, description, url string) (entity.Video, error)
	DeleteById(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id, title, description, url string) error
}
