package usecases

import (
	"errors"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type SateliteFinder struct {
	repository domain.SateliteRepository
}

func NewSateliteFinder(repository domain.SateliteRepository) SateliteFinder {
	return SateliteFinder{
		repository: repository,
	}
}

func (s SateliteFinder) Find(name string) (domain.Satelite, error) {
	satelite, err := s.repository.FindByName(name)

	if err != nil {
		return domain.Satelite{}, errors.New("satelite " + name + " not found")
	}

	return satelite, nil
}
