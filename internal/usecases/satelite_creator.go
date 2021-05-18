package usecases

import (
	core "github.com/franciscoruizar/quasar-fire/internal"
	"github.com/google/uuid"
)

type SateliteCreator struct {
	repository core.SateliteRepository
}

func NewSateliteCreator(repository core.SateliteRepository) SateliteCreator {
	return SateliteCreator{
		repository: repository,
	}
}

func (s SateliteCreator) Create(name string, x float64, y float64, distance float64, message []string) (core.Satelite, error) {
	id := uuid.NewString()
	satelite, err := core.NewSatelite(id, name, x, y, distance, message)
	if err != nil {
		return satelite, err
	}
	error := s.repository.Save(satelite)

	return satelite, error
}
