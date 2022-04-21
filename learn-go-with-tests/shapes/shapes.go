package shapes

import "math"

type Rectangle struct {
	W, H float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.W + rectangle.H)
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
