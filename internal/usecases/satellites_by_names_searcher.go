package usecases

import (
	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type SatellitesByNamesSearcher struct {
	finder SateliteFinder
}

func NewSatellitesByNamesSearcher(repository domain.SateliteRepository) SatellitesByNamesSearcher {
	return SatellitesByNamesSearcher{
		finder: NewSateliteFinder(repository),
	}
}

func (s SatellitesByNamesSearcher) Search(names []string) ([]domain.Satelite, error) {
	var values []domain.Satelite

	for i := 0; i < len(names); i++ {
		satelite, err := s.finder.Find(names[i])
		if err != nil {
			return nil, err
		}

		values = append(values, satelite)
	}

	return values, nil
}
