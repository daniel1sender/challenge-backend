package http

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	video, err := h.useCase.GetByID(c.Request.Context(), id)
	if err != nil {
		var responseError Error
		switch {
		case errors.Is(err, videos.ErrNoVideoFound):
			responseError.Error = err.Error()
			c.JSON(http.StatusNotFound, responseError)
		default:
			responseError.Error = err.Error()
			c.JSON(http.StatusInternalServerError, responseError)
		}
		return
	}
	c.JSON(http.StatusOK, video)
}
