package usecases

import (
	core "github.com/franciscoruizar/quasar-fire/internal"
)

type LocationFinder struct {
	repository core.SateliteRepository
}

func NewLocationFinder(repository core.SateliteRepository) LocationFinder {
	return LocationFinder{
		repository: repository,
	}
}

func (finder LocationFinder) Find(distances []float64) (core.Position, error) {
	values := finder.repository.FindAll()

	if len(distances) < len(values) {
		return core.Position{}, nil
	}

	var circles []core.Circle
	for i := 0; i < len(values); i++ {
		circles = append(circles, core.NewCircle(values[i].Position.X, values[i].Position.Y, distances[i]))
	}

	if len(circles) < 2 {
		return core.Position{}, nil
	}

	return core.Position{}, nil
}
