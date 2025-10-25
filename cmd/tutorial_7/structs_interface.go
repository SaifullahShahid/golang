package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Length float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	height float64
	base   float64
}
type Measurable interface {
	Perimeter() float64
}
type Geometry interface {
	Shape
	Measurable
}

func (t Triangle) Area() float64 {
	return t.height * t.base / 2
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func describeShape(g Geometry) (float64, float64) {
	return g.Area(), g.Perimeter()
}
func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}

type CalculationError struct {
	message string
}

func (ce CalculationError) Error() string {
	return ce.message
}
func squareRoot(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{message: "Argument must be a non-negative number"}
	}
	return math.Sqrt(val), nil
}
