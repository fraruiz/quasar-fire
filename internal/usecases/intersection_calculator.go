package usecases

import (
	"errors"
	"math"

	core "github.com/franciscoruizar/quasar-fire/internal"
)

type IntersectionCalculator struct {
}

func NewIntersectionCalculator() IntersectionCalculator {
	return IntersectionCalculator{}
}

func (c IntersectionCalculator) Calculate(a core.Circle, b core.Circle) ([]core.Position, error) {
	p := []core.Position{}

	dx, dy := b.Center.X-a.Center.X, b.Center.Y-a.Center.Y
	lr := a.Radius + b.Radius           //radius and
	dr := math.Abs(a.Radius - b.Radius) //radius difference
	ab := math.Sqrt(dx*dx + dy*dy)      //center distance

	if ab > lr && ab < dr {
		return p, errors.New("not intersects")
	}

	theta1 := math.Atan(dy / dx)
	ef := lr - ab
	ao := a.Radius - ef/2
	theta2 := math.Acos(ao / a.Radius)
	theta := theta1 + theta2
	xc := a.Center.X + a.Radius*math.Cos(theta)
	yc := a.Center.Y + a.Radius*math.Sin(theta)
	p = append(p, core.NewPosition(xc, yc))
	if ab < lr {
		theta3 := math.Acos(ao / a.Radius)
		theta = theta3 - theta1
		xd := a.Center.X + a.Radius*math.Cos(theta)
		yd := a.Center.Y - a.Radius*math.Sin(theta)
		p = append(p, core.NewPosition(xd, yd))
	}
	return p, nil
}
