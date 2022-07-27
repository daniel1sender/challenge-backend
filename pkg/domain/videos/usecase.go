package videos

import (
	"context"
	"errors"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

var (
	ErrNoVideoFound = errors.New("no video found")
)

type UseCase interface {
	GetAll(ctx context.Context) ([]entities.Video, error)
	GetByID(ctx context.Context, id string) (entities.Video, error)
	Create(ctx context.Context, title, description, url string) (entities.Video, error)
	DeleteById(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id, title, description, url string) error
}
