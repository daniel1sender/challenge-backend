package usecases

import "github.com/daniel1sender/alura-flix/pkg/domain/videos"

type VideoUseCase struct {
	storage videos.Repository
}

func NewUseCase(storage videos.Repository) VideoUseCase {
	return VideoUseCase{storage: storage}
}
