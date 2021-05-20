package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type SateliteID struct {
	value string
}

func (id SateliteID) Value() string {
	return id.value
}

var ErrInvalidSateliteID = errors.New("invalid  ID")

func NewSateliteID(value string) (SateliteID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return SateliteID{}, fmt.Errorf("%w: %s", ErrInvalidSateliteID, value)
	}

	return SateliteID{
		value: v.String(),
	}, nil
}

var ErrEmptySateliteName = errors.New("the field Name can not be empty")

type SateliteName struct {
	value string
}

func NewSateliteName(value string) (SateliteName, error) {
	if value == "" {
		return SateliteName{}, ErrEmptySateliteName
	}

	return SateliteName{
		value: value,
	}, nil
}

func (name SateliteName) Value() string {
	return name.value
}

type Satelite struct {
	id       SateliteID
	name     SateliteName
	position Position
	distance float64
	message  []string
}

func NewSatelite(id string, name string, x float64, y float64, distance float64, message []string) (Satelite, error) {
	idVO, err := NewSateliteID(id)

	if err != nil {
		return Satelite{}, err
	}

	nameVO, err := NewSateliteName(name)

	if err != nil {
		return Satelite{}, err
	}

	positionVO, err := NewPosition(x, y)

	if err != nil {
		return Satelite{}, err
	}

	return Satelite{
		id:       idVO,
		name:     nameVO,
		position: positionVO,
		distance: distance,
		message:  message,
	}, nil
}

func (s Satelite) ID() SateliteID {
	return s.id
}

func (s Satelite) Name() SateliteName {
	return s.name
}

func (s Satelite) Position() Position {
	return s.position
}

func (s Satelite) Distance() float64 {
	return s.distance
}

func (s Satelite) Message() []string {
	return s.message
}

type SateliteRepository interface {
	Update(satelite Satelite) error
	FindAll() ([]Satelite, error)
	FindByName(name string) (Satelite, error)
}
