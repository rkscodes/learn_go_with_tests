package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Raidus float64
}

type Triangle struct {
	Length, Width float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Raidus * c.Raidus
}

func (t Triangle) Area() float64 {
	return t.Length * t.Width * 0.5
}
