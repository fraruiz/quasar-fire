package usecases

import (
	"errors"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type MessageUpdater struct {
	repository domain.SateliteRepository
}

func NewMessageUpdater(repository domain.SateliteRepository) MessageUpdater {
	return MessageUpdater{
		repository: repository,
	}
}

func (updater MessageUpdater) Update(name string, distance float64, message []string) error {
	satelite, err := updater.repository.FindByName(name)

	if err != nil {
		return errors.New("satelite " + name + " not found")
	}

	sateliteUpdated, err := domain.NewSatelite(satelite.ID().Value(), name, satelite.Position().X().Value(), satelite.Position().Y().Value(), distance, message)

	if err != nil {
		return err
	}

	updater.repository.Update(sateliteUpdated)

	return nil
}
