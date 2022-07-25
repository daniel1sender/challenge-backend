package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entity"
)

func (u VideoUseCase) GetByID(ctx context.Context, id string) (entity.Video, error) {
	return u.storage.GetByID(ctx, id)
}
