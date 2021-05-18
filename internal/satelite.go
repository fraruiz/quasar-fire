package core

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type SateliteID struct {
	value string
}

var ErrInvalidSateliteID = errors.New("invalid Course ID")

func NewSateliteID(value string) (SateliteID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return SateliteID{}, fmt.Errorf("%w: %s", ErrInvalidSateliteID, value)
	}

	return SateliteID{
		value: v.String(),
	}, nil
}

type Satelite struct {
	ID       SateliteID
	Name     string
	Position Position
	Distance float64
	Message  []string
}

func NewSatelite(id string, name string, x float64, y float64, distance float64, message []string) (Satelite, error) {
	idVO, err := NewSateliteID(id)

	if err != nil {
		return Satelite{}, err
	}

	return Satelite{
		ID:       idVO,
		Name:     name,
		Position: NewPosition(x, y),
		Distance: distance,
		Message:  message,
	}, nil
}

type SateliteRepository interface {
	Save(satelite Satelite) error
	FindAll() []Satelite
	FindByName(name string) (Satelite, error)
}
