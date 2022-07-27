package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u CategoryUseCase) GetByID(ctx context.Context, id string) (entities.Category, error) {
	return u.storage.GetByID(ctx, id)
}
