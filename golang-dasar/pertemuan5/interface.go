package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Parameter() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Parameter() float64 {
	return 2 * r.Width + 2 * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Parameter() float64 {
	return 2 * math.Pi * c.Radius
}

func main ()  {
	var s Shape

	s = Rectangle {Width: 5, Height: 8}
	fmt.Println("Rectangle Area = ", s.Area())
	fmt.Println("Rectangle Parameter = ", s.Parameter())

	s = Circle{Radius: 4}
	
	fmt.Println("Circle area = ", s.Area())
	fmt.Println("Circle parameter = ", s.Parameter())
}