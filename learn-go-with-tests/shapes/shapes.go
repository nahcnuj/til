package shapes

import "math"

type Shape interface {
	Area() float64
}

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

type Triangle struct {
	// 底辺
	Base float64

	// 高さ
	Height float64
}

func (t Triangle) Area() float64 {
	return 0
}
