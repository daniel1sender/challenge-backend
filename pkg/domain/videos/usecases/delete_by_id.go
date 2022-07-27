package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos"
)

func (u VideoUseCase) DeleteById(ctx context.Context, id string) error {
	_, err := u.storage.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("could not delete the video: %w", videos.ErrNoVideoFound)
	}
	return u.storage.DeleteById(ctx, id)
}
