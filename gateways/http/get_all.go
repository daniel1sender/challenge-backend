package http

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetAll(c *gin.Context) {

	listOfVideo, err := h.useCase.GetAll(c.Request.Context())
	if err != nil {
		var responseError Error
		switch {
		case errors.Is(err, domain.ErrNoVideoFound):
			responseError.Error = err.Error()
			c.JSON(http.StatusNotFound, responseError)
		default:
			responseError.Error = err.Error()
			c.JSON(http.StatusInternalServerError, responseError)
		}
		return
	}
	c.JSON(http.StatusOK, listOfVideo)
}
