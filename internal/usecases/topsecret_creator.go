package usecases

import (
	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	dto "github.com/franciscoruizar/quasar-fire/internal/usecases/dto"
)

type TopSecretCreator struct {
	messageDecoder MessageDecoder
	locationFinder LocationFinder
}

func NewTopSecretCreator(repository domain.SateliteRepository) TopSecretCreator {
	return TopSecretCreator{
		messageDecoder: NewMessageDecoder(),
		locationFinder: NewLocationFinder(repository),
	}
}

type TopSecretCreatorRequest struct {
	Name      string
	Dinstance float64
	Message   []string
}

type TopSecretCreatorResponse struct {
	Position dto.PositionResponse
	Message  string
}

func (creator TopSecretCreator) Create(requests []TopSecretCreatorRequest) (TopSecretCreatorResponse, error) {

	var distances []float64
	var messages [][]string
	var satellites []string

	for i := 0; i < len(requests); i++ {
		request := requests[i]
		messages = append(messages, request.Message)
		distances = append(distances, request.Dinstance)
		satellites = append(satellites, request.Name)
	}

	message := creator.messageDecoder.Decode(messages)
	position, err := creator.locationFinder.Find(distances, satellites)

	if err != nil {
		return TopSecretCreatorResponse{}, err
	}

	return TopSecretCreatorResponse{
		Position: dto.NewPositionResponse(position.X().Value(), position.Y().Value()),
		Message:  message,
	}, nil
}
