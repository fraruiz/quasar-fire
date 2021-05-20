package topsecretsplit

import (
	"errors"
	"net/http"

	"github.com/franciscoruizar/quasar-fire/internal/domain"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/server/handler"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/franciscoruizar/quasar-fire/internal/usecases/dto"
	"github.com/gin-gonic/gin"
)

type TopSecretSplitCreateRequest struct {
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type TopSecretSplitCreateResponse struct {
	Position dto.PositionResponse `json:"position" binding:"required"`
	Message  string               `json:"message" binding:"required"`
}

func TopSecretSplitHandler(service usecases.TopSecretSplitCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TopSecretSplitCreateRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		sateliteName := c.Params.ByName("name")

		response, err := service.Create(sateliteName, req.Distance, req.Message)

		if err != nil {
			errorMessage := handler.Error{
				Message: err.Error(),
			}
			switch {
			case errors.Is(err, domain.ErrInvalidSateliteID):
				c.JSON(http.StatusBadRequest, errorMessage)
				return
			default:
				c.JSON(http.StatusInternalServerError, errorMessage)
				return
			}
		}

		responseMap := TopSecretSplitCreateResponse{
			Message:  response.Message,
			Position: response.Position,
		}

		c.Status(http.StatusCreated)

		c.JSON(http.StatusCreated, responseMap)
	}
}
