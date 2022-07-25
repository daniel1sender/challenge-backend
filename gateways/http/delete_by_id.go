package http

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := h.useCase.DeleteById(c.Request.Context(), id)
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
	c.Status(http.StatusOK)
}
