package videos

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos"
	server "github.com/daniel1sender/alura-flix/pkg/gateways/http"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetAll(c *gin.Context) {

	listOfVideo, err := h.useCase.GetAll(c.Request.Context())
	if err != nil {
		var responseError server.Error
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
	c.JSON(http.StatusOK, listOfVideo)
}
