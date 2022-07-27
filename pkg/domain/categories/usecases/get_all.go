package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u CategoryUseCase) GetAll(ctx context.Context) ([]entities.Category, error) {
	return u.storage.GetAll(ctx)
}
