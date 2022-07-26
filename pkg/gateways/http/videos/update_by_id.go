package videos

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos"
	server "github.com/daniel1sender/alura-flix/pkg/gateways/http"
	"github.com/gin-gonic/gin"
)

type UpdateRequestVideo struct {
	Title       string `form:"Title"`
	Description string `form:"Description"`
	URL         string `form:"URL"`
}

type UpdateResponseVideo struct {
	Title       string `form:"Title"`
	Description string `form:"Description"`
	URL         string `form:"URL"`
}

func (h Handler) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var video UpdateRequestVideo
	if err := c.ShouldBindJSON(&video); err != nil {
		var responseError server.Error
		responseError.Error = err.Error()
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	if len(video.Title) == 0 && len(video.Description) == 0 && len(video.URL) == 0 {
		var responseError server.Error
		responseError.Error = server.ErrEmptyBody.Error()
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	err := h.useCase.UpdateByID(c.Request.Context(), id, video.Title, video.Description, video.URL)
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

	newVideo, _ := h.useCase.GetByID(c.Request.Context(), id)
	c.JSON(http.StatusOK, newVideo)
}
