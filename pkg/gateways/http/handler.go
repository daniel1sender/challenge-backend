package http

import "github.com/daniel1sender/alura-flix/pkg/domain"

type Handler struct {
	useCase domain.UseCase
}

func NewHandler(usecase domain.UseCase) Handler {
	return Handler{
		useCase: usecase,
	}
}
