package domain

import (
	"errors"
	"math"
)

var ErrInvalidPoint = errors.New("invalid point")

type Point struct {
	value float64
}

func NewPoint(value float64) (Point, error) {
	if math.IsNaN(value) {
		return Point{}, ErrInvalidPoint
	}

	if value == -0 {
		return Point{value: 0}, nil
	}

	return Point{value: value}, nil
}

func (p Point) Value() float64 {
	return p.value
}
