package memory

import (
	"errors"

	core "github.com/franciscoruizar/quasar-fire/internal"
	"github.com/google/uuid"
)

type SateliteRepository struct {
	values []core.Satelite
}

func NewSateliteRepository() *SateliteRepository {
	kenobi, err := core.NewSatelite(uuid.NewString(), "kenobi", -500, -200)
	skywalker, err1 := core.NewSatelite(uuid.NewString(), "skywalker", 100, -100)
	sato, err2 := core.NewSatelite(uuid.NewString(), "sato", 500, 100)

	if err != nil || err1 != nil || err2 != nil {
		return &SateliteRepository{
			values: []core.Satelite{},
		}
	}

	return &SateliteRepository{
		values: []core.Satelite{kenobi, skywalker, sato},
	}
}

func (r *SateliteRepository) Save(satelite core.Satelite) error {
	return nil
}

func (r *SateliteRepository) FindAll() ([]core.Satelite, error) {
	return r.values, nil
}

func (r *SateliteRepository) FindByName(name string) (core.Satelite, error) {
	values := r.values

	for i := 0; i < len(values); i++ {
		if values[i].Name == name {
			return values[i], nil
		}
	}

	return core.Satelite{}, errors.New("not found")
}
