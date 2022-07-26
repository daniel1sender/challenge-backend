package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

func (u VideoUseCase) Create(ctx context.Context, title, description, url string) (entities.Video, error) {
	newVideo, err := entities.NewVideo(title, description, url)
	if err != nil {
		return entities.Video{}, fmt.Errorf("could not create the video due to: %w", err)
	}
	err = u.storage.Insert(ctx, newVideo)
	if err != nil {
		return entities.Video{}, fmt.Errorf("could not insert the video in the database due to: %w", err)
	}
	return newVideo, nil
}
