package core

type Circle struct {
	Center Position
	Radius float64
}

func NewCircle(x, y, r float64) Circle {
	point := NewPosition(x, y)
	return Circle{point, r}
}
