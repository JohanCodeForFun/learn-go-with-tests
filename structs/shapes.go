package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Heigth float64
}

type Triangle struct {
	Base   float64
	Heigth float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Heigth) * 0.5
}

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Heigth + rectangle.Width)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Heigth * rectangle.Width
// }
