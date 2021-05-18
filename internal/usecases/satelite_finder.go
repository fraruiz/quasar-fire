package usecases

import (
	"errors"

	core "github.com/franciscoruizar/quasar-fire/internal"
)

type SateliteFinder struct {
	repository core.SateliteRepository
}

func NewSateliteFinder(repository core.SateliteRepository) SateliteFinder {
	return SateliteFinder{
		repository: repository,
	}
}

func (s SateliteFinder) Find(name string) (core.Satelite, error) {
	satelite, err := s.repository.FindByName(name)

	if err != nil {
		return core.Satelite{}, errors.New("satelite " + name + " not found")
	}

	return satelite, nil
}
