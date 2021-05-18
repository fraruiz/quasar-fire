package usecases

import (
	core "github.com/franciscoruizar/quasar-fire/internal"
)

type TopSecretCreator struct {
	sateliteCreator SateliteCreator
	decoder         MessageDecoder
	finder          LocationFinder
}

func NewTopSecretCreator(repository core.SateliteRepository) TopSecretCreator {
	return TopSecretCreator{
		sateliteCreator: NewSateliteCreator(repository),
		decoder:         NewMessageDecoder(),
		finder:          NewLocationFinder(repository),
	}
}

type TopSecretCreatorRequest struct {
	Name      string
	Dinstance float64
	Message   []string
}

type TopSecretCreatorResponse struct {
	Position core.Position
	Message  string
}

func (creator TopSecretCreator) Create(requests []TopSecretCreatorRequest) (TopSecretCreatorResponse, error) {

	var distances []float64
	var messages [][]string

	for i := 0; i < len(requests); i++ {
		request := requests[i]
		messages = append(messages, request.Message)
		distances = append(distances, request.Dinstance)
	}

	message := creator.decoder.Decode(messages)
	position, err := creator.finder.Find(distances)

	if err != nil {
		return TopSecretCreatorResponse{}, err
	}

	return TopSecretCreatorResponse{
		Position: position,
		Message:  message,
	}, nil
}
