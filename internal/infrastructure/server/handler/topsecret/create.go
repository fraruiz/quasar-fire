package topsecret

import (
	"errors"
	"net/http"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/server/handler"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/franciscoruizar/quasar-fire/internal/usecases/dto"
	"github.com/gin-gonic/gin"
)

type TopSecretCreateRequests struct {
	Satellites []TopSecretCreateRequest `json:"satellites" binding:"required"`
}

type TopSecretCreateRequest struct {
	Name     string   `json:"name" binding:"required"`
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type TopSecretCreateResponse struct {
	Position dto.PositionResponse `json:"position" binding:"required"`
	Message  string               `json:"message" binding:"required"`
}

func TopSecretHandler(service usecases.TopSecretCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TopSecretCreateRequests
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		var satellites []dto.TopSecretRequest
		for i := 0; i < len(req.Satellites); i++ {
			satellites = append(satellites, dto.TopSecretRequest{
				Name:     req.Satellites[i].Name,
				Distance: req.Satellites[i].Distance,
				Message:  req.Satellites[i].Message,
			})
		}

		response, err := service.Create(satellites)

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

		responseMap := TopSecretCreateResponse{
			Message:  response.Message,
			Position: response.Position,
		}

		c.Status(http.StatusCreated)

		c.JSON(http.StatusCreated, responseMap)
	}
}
