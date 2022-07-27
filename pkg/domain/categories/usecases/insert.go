package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u CategoryUseCase) Insert(ctx context.Context, category entities.Category) error {
	return u.storage.Insert(ctx, category)
}
