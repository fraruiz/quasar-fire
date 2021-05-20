package domain

type Circle struct {
	center Position
	radius Point
}

func NewCircle(x, y, radius float64) (Circle, error) {
	centerVO, err := NewPosition(x, y)

	if err != nil {
		return Circle{}, err
	}

	radiusVO, err := NewPoint(radius)

	if err != nil {
		return Circle{}, err
	}

	return Circle{centerVO, radiusVO}, nil
}

func (c Circle) Radius() Point {
	return c.radius
}

func (c Circle) Center() Position {
	return c.center
}
