package shapes

type Rectangle struct {
	w, h float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.w + rectangle.h)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.w * rectangle.h
}
