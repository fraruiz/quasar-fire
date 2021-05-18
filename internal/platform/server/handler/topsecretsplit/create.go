package topsecretsplit

import (
	"errors"
	"net/http"

	core "github.com/franciscoruizar/quasar-fire/internal"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

func TopSecretSplitCreateHandler(service usecases.TopSecretSplitCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		err := service.Create(c.Params.ByName("name"), req.Distance, req.Message)

		if err != nil {
			switch {
			case errors.Is(err, core.ErrInvalidSateliteID):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}
