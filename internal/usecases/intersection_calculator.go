package usecases

import (
	"errors"
	"math"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type IntersectionCalculator struct {
	distanceCalculator DistanceCalculator
}

func NewIntersectionCalculator() IntersectionCalculator {
	return IntersectionCalculator{
		distanceCalculator: NewDistanceCalculator(),
	}
}

func (c IntersectionCalculator) Calculate(circle0 domain.Circle, circle1 domain.Circle) ([]domain.Position, error) {
	distance := c.distanceCalculator.Calculate(circle0.Center(), circle1.Center())

	err := c.ensureDistanceAndCirclesData(distance, circle0, circle1)

	if err != nil {
		return nil, err
	}

	a := (math.Pow(circle0.Radius().Value(), 2) - math.Pow(circle1.Radius().Value(), 2) + math.Pow(distance, 2)) / (2 * distance)

	h := math.Sqrt(math.Pow(circle0.Radius().Value(), 2) - math.Pow(a, 2))

	x2 := circle0.Center().X().Value() + a*(circle1.Center().X().Value()-circle0.Center().X().Value())/distance
	y2 := circle0.Center().Y().Value() + a*(circle1.Center().Y().Value()-circle0.Center().Y().Value())/distance

	x3 := x2 + h*(circle1.Center().Y().Value()-circle0.Center().Y().Value())/distance
	x3 = math.Round(x3)
	y3 := y2 - h*(circle1.Center().X().Value()-circle0.Center().X().Value())/distance
	y3 = math.Round(y3)

	position0, err := domain.NewPosition(x3, y3)
	if err != nil {
		return nil, err
	}

	x4 := x2 - h*(circle1.Center().Y().Value()-circle0.Center().Y().Value())/distance
	x4 = math.Round(x4)
	y4 := y2 + h*(circle1.Center().X().Value()-circle0.Center().X().Value())/distance
	y4 = math.Round(y4)

	position1, err := domain.NewPosition(x4, y4)
	if err != nil {
		return nil, err
	}

	return []domain.Position{position0, position1}, nil
}

func (c IntersectionCalculator) ensureDistanceAndCirclesData(distance float64, a domain.Circle, b domain.Circle) error {
	if distance > a.Radius().Value()+b.Radius().Value() {
		return errors.New("non intersecting between circles")
	}

	radiusSubtraction := a.Radius().Value() - b.Radius().Value()

	if distance < math.Abs(radiusSubtraction) {
		return errors.New("one circle within other")
	}

	if distance == 0 && a.Radius().Value() == b.Radius().Value() {
		return errors.New("coincident circles")
	}

	return nil
}
