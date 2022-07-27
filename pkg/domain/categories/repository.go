package categories

import (
	"context"
	"errors"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

var (
	ErrNoCategoryFound = errors.New("no category found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]entities.Category, error)
	Insert(ctx context.Context, category entities.Category) error
	GetByID(ctx context.Context, id string) (entities.Category, error)
	DeleteById(ctx context.Context, id string) error
	UpdateTitleByID(ctx context.Context, id, title string) error
	UpdateColorByID(ctx context.Context, id, color string) error
}
