package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u VideoUseCase) GetByID(ctx context.Context, id string) (entities.Video, error) {
	return u.storage.GetByID(ctx, id)
}
