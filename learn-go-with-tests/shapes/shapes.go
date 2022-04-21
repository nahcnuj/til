package shapes

type Rectangle struct {
	W, H float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.W + rectangle.H)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.W * rectangle.H
}

type Circle struct {
	Radius float64
}
