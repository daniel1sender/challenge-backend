package videos

import "github.com/daniel1sender/alura-flix/pkg/domain/videos"

type Handler struct {
	useCase videos.UseCase
}

func NewHandler(usecase videos.UseCase) Handler {
	return Handler{
		useCase: usecase,
	}
}
