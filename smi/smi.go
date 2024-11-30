package smi

import "math"

func Perimeter(p Rectangle) float64 {
	return (p.Width + p.Height) * 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base / 2) * t.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}
