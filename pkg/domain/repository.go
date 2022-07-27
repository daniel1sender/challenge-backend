package domain

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entities.Video, error)
	Insert(ctx context.Context, video entities.Video) error
	GetByID(ctx context.Context, id string) (entities.Video, error)
	DeleteById(ctx context.Context, id string) error
	UpdateTitleByID(ctx context.Context, id, title string) error
	UpdateURLByID(ctx context.Context, id, url string) error
	UpdateDescriptionByID(ctx context.Context, id, description string) error
}
