package http

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/domain/validate"
	"github.com/gin-gonic/gin"
)

type CreateRequestVideo struct {
	Title       string `form:"Title"`
	Description string `form:"Description"`
	URL         string `form:"URL"`
}

func (h Handler) Create(c *gin.Context) {

	var video CreateRequestVideo

	if err := c.ShouldBindJSON(&video); err != nil {
		var responseError Error
		responseError.Error = err.Error()
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	if len(video.Title) == 0 && len(video.Description) == 0 && len(video.URL) == 0 {
		var responseError Error
		responseError.Error = ErrEmptyBody.Error()
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	newVideo, err := h.useCase.Create(c.Request.Context(), video.Title, video.Description, video.URL)
	if err != nil {
		var responseError Error
		switch {
		case errors.Is(err, validate.ErrEmptyTitle):
			responseError.Error = err.Error()
			c.JSON(http.StatusBadRequest, responseError)

		case errors.Is(err, validate.ErrInvalidTitle):
			responseError.Error = err.Error()
			c.JSON(http.StatusBadRequest, responseError)

		case errors.Is(err, validate.ErrEmptyDescription):
			responseError.Error = err.Error()
			c.JSON(http.StatusBadRequest, responseError)

		case errors.Is(err, validate.ErrInvalidDescription):
			responseError.Error = err.Error()
			c.JSON(http.StatusBadRequest, responseError)

		case errors.Is(err, validate.ErrEmptyURL):
			responseError.Error = err.Error()
			c.JSON(http.StatusBadRequest, responseError)

		default:
			responseError.Error = err.Error()
			c.JSON(http.StatusInternalServerError, responseError)
		}
		return
	}
	c.JSON(http.StatusCreated, newVideo)
}
