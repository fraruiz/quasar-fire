package usecases

import (
	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	dto "github.com/franciscoruizar/quasar-fire/internal/usecases/dto"
)

type TopSecretCreator struct {
	messageDecoder MessageDecoder
	locationFinder LocationFinder
	sateliteFinder SateliteFinder
}

func NewTopSecretCreator(sateliteRepository domain.SateliteRepository) TopSecretCreator {
	return TopSecretCreator{
		messageDecoder: NewMessageDecoder(),
		locationFinder: NewLocationFinder(sateliteRepository),
		sateliteFinder: NewSateliteFinder(sateliteRepository),
	}
}

func (creator TopSecretCreator) Create(requests []dto.TopSecretRequest) (dto.TopSecretResponse, error) {
	var distances []float64
	var messages []string
	var satellitesName []string

	for i := 0; i < len(requests); i++ {
		request := requests[i]
		messages = append(messages, request.Message...)
		distances = append(distances, request.Distance)
		satellitesName = append(satellitesName, request.Name)
	}

	message, err := creator.messageDecoder.Decode(messages)

	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	position, err := creator.locationFinder.Find(distances, satellitesName)

	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	return dto.TopSecretResponse{
		Position: dto.NewPositionResponse(position.X().Value(), position.Y().Value()),
		Message:  message,
	}, nil
}
