package usecases

import "github.com/daniel1sender/alura-flix/pkg/domain"

type VideoUseCase struct {
	storage domain.Repository
}

func NewUseCase(storage domain.Repository) VideoUseCase {
	return VideoUseCase{storage: storage}
}
