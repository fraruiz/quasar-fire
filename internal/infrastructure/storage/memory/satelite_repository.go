package memory

import (
	"errors"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	"github.com/google/uuid"
)

type SateliteRepository struct {
	values []domain.Satelite
}

func NewSateliteRepository() *SateliteRepository {
	kenobi, err := domain.NewSatelite(uuid.NewString(), "kenobi", -500, -200, 0, nil)
	skywalker, err1 := domain.NewSatelite(uuid.NewString(), "skywalker", 100, -100, 0, nil)
	sato, err2 := domain.NewSatelite(uuid.NewString(), "sato", 500, 100, 0, nil)

	if err != nil || err1 != nil || err2 != nil {
		return &SateliteRepository{
			values: []domain.Satelite{},
		}
	}

	return &SateliteRepository{
		values: []domain.Satelite{kenobi, skywalker, sato},
	}
}

func (r *SateliteRepository) Update(satelite domain.Satelite) error {
	for i := 0; i < len(r.values); i++ {
		if r.values[i].ID().Value() == satelite.ID().Value() {
			r.values[i] = satelite
		}
	}

	return nil
}

func (r *SateliteRepository) FindAll() ([]domain.Satelite, error) {
	return r.values, nil
}

func (r *SateliteRepository) FindByName(name string) (domain.Satelite, error) {
	values := r.values

	for i := 0; i < len(values); i++ {
		if values[i].Name().Value() == name {
			return values[i], nil
		}
	}

	return domain.Satelite{}, errors.New("not found")
}
