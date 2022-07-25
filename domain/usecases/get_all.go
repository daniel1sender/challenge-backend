package usecases

import (
	"context"

	"github.com/daniel1sender/alura-flix/domain/entity"
)

func (u VideoUseCase) GetAll(ctx context.Context) ([]entity.Video, error) {
	return u.storage.GetAll(ctx)
}
