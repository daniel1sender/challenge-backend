package usecases

import "github.com/daniel1sender/alura-flix/domain"

type VideoUseCase struct {
	storage domain.Repository
}

func NewUseCase(storage domain.Repository) VideoUseCase {
	return VideoUseCase{storage: storage}
}
