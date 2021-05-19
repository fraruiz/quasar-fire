package domain

type Position struct {
	x, y Point
}

func NewPosition(x, y float64) (Position, error) {
	xVO, err := NewPoint(x)

	if err != nil {
		return Position{}, err
	}

	yVO, err := NewPoint(y)

	if err != nil {
		return Position{}, err
	}

	return Position{xVO, yVO}, nil
}

func (p Position) X() Point {
	return p.x
}

func (p Position) Y() Point {
	return p.y
}
