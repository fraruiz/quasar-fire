package topsecretsplit

import (
	"errors"
	"net/http"

	core "github.com/franciscoruizar/quasar-fire/internal"
	"github.com/franciscoruizar/quasar-fire/internal/platform/server/handler"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/gin-gonic/gin"
)

type TopSecretSplitFindResponse struct {
	Name     string                   `json:"name" binding:"required"`
	Position handler.PositionResponse `json:"position" binding:"required"`
	Distance float64                  `json:"distance" binding:"required"`
	Message  []string                 `json:"message" binding:"required"`
}

func TopSecretSplitFindHandler(service usecases.SateliteFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Params.ByName("name")
		response, err := service.Find(param)

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

		c.Status(http.StatusAccepted)

		c.JSON(http.StatusAccepted, TopSecretSplitFindResponse{
			Name:     response.Name,
			Position: handler.NewPositionResponse(response.Position.X, response.Position.Y),
			Distance: response.Distance,
			Message:  response.Message,
		})
	}
}
