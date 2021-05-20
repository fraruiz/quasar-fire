package usecases

import (
	"math"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type DistanceCalculator struct {
}

func NewDistanceCalculator() DistanceCalculator {
	return DistanceCalculator{}
}

func (c DistanceCalculator) Calculate(a domain.Position, b domain.Position) float64 {
	x := math.Pow(b.X().Value()-a.X().Value(), 2)
	y := math.Pow(b.Y().Value()-a.Y().Value(), 2)

	return math.Sqrt(x + y)
}
