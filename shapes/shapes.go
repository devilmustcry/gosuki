package shapes

import (
	"math"
)

// Rectangle : Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle : Circle struct
type Circle struct {
	Radius float64
}

// Perimeter : Find perimeter
func (r Rectangle) Perimeter() float64 {
	return 2.0 * (r.Width + r.Height)
}

// Perimeter : Find perimeter
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter : Find perimeter
func (c Circle) Perimeter() float64 {
	return math.Pi * 2.0 * c.Radius
}

// Area : Find area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}
