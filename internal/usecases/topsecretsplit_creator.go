package usecases

import (
	"errors"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	dto "github.com/franciscoruizar/quasar-fire/internal/usecases/dto"
)

type TopSecretSplitCreator struct {
	messageUpdater     MessageUpdater
	satellitesSearcher SatellitesSearcher
	messageDecoder     MessageDecoder
	locationFinder     LocationFinder
}

func NewTopSecretSplitCreator(sateliteRepository domain.SateliteRepository) TopSecretSplitCreator {
	return TopSecretSplitCreator{
		messageUpdater:     NewMessageUpdater(sateliteRepository),
		satellitesSearcher: NewSatellitesSearcher(sateliteRepository),
		messageDecoder:     NewMessageDecoder(),
		locationFinder:     NewLocationFinder(sateliteRepository),
	}
}

func (creator TopSecretSplitCreator) Create(name string, distance float64, message []string) (dto.TopSecretResponse, error) {
	err := creator.messageUpdater.Update(name, distance, message)

	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	satelites, err := creator.satellitesSearcher.Search()

	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	err = creator.ensureSatellites(satelites)

	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	var distances []float64
	var messages []string
	var satellitesName []string

	for i := 0; i < len(satelites); i++ {
		satelite := satelites[i]
		messages = append(messages, satelite.Message()...)
		distances = append(distances, satelite.Distance())
		satellitesName = append(satellitesName, satelite.Name().Value())
	}

	position, err := creator.locationFinder.Find(distances, satellitesName)
	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	messageDecode, err := creator.messageDecoder.Decode(messages)
	if err != nil {
		return dto.TopSecretResponse{}, err
	}

	return dto.TopSecretResponse{
		Position: dto.NewPositionResponse(position.X().Value(), position.Y().Value()),
		Message:  messageDecode,
	}, nil
}

func (creator TopSecretSplitCreator) ensureSatellites(values []domain.Satelite) error {
	for _, item := range values {
		if item.Distance() == 0 {
			return errors.New(item.Name().Value() + " distance is zero")
		}

		if item.Message() == nil {
			return errors.New("the " + item.Name().Value() + " messagge is zero")
		}
	}

	return nil

}
