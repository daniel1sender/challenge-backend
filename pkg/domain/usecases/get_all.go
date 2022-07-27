package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u VideoUseCase) GetAll(ctx context.Context) ([]entities.Video, error) {
	return u.storage.GetAll(ctx)
}
