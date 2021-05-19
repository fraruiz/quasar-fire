package usecases

import (
	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type SatellitesSearcher struct {
	repository domain.SateliteRepository
}

func NewSatellitesSearcher(repository domain.SateliteRepository) SatellitesSearcher {
	return SatellitesSearcher{
		repository: repository,
	}
}

func (s SatellitesSearcher) Search() ([]domain.Satelite, error) {

	values, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return values, nil
}
