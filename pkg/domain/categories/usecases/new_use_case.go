package usecases

import "github.com/daniel1sender/alura-flix/pkg/domain/categories"

type CategoryUseCase struct {
	storage categories.Repository
}

func NewCategoryUseCase(storage categories.Repository) CategoryUseCase {
	return CategoryUseCase{storage: storage}
}
