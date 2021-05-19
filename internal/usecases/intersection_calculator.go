package usecases

import (
	"errors"
	"math"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type IntersectionCalculator struct {
}

func NewIntersectionCalculator() IntersectionCalculator {
	return IntersectionCalculator{}
}

func (c IntersectionCalculator) Calculate(a domain.Circle, b domain.Circle) ([]domain.Position, error) {
	p := []domain.Position{}

	dx, dy := b.Center().X().Value()-a.Center().X().Value(), b.Center().Y().Value()-a.Center().Y().Value()
	lr := a.Radius().Value() + b.Radius().Value()           //radius and
	dr := math.Abs(a.Radius().Value() - b.Radius().Value()) //radius difference
	ab := math.Sqrt(dx*dx + dy*dy)                          //center distance

	if ab > lr && ab < dr {
		return p, errors.New("not intersects")
	}

	theta1 := math.Atan(dy / dx)
	ef := lr - ab
	ao := a.Radius().Value() - ef/2
	theta2 := math.Acos(ao / a.Radius().Value())
	theta := theta1 + theta2
	xc := a.Center().X().Value() + a.Radius().Value()*math.Cos(theta)
	yc := a.Center().Y().Value() + a.Radius().Value()*math.Sin(theta)

	position, err := domain.NewPosition(xc, yc)

	if err != nil {
		return nil, err
	}

	p = append(p, position)
	if ab < lr {
		theta3 := math.Acos(ao / a.Radius().Value())
		theta = theta3 - theta1
		xd := a.Center().X().Value() + a.Radius().Value()*math.Cos(theta)
		yd := a.Center().Y().Value() - a.Radius().Value()*math.Sin(theta)

		position, err = domain.NewPosition(xd, yd)

		if err != nil {
			return nil, err
		}

		p = append(p, position)
	}
	return p, nil
}
