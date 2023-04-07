package main

import (
	"fmt"
	"math"
)

func main()  {
	a := 30.0

	squareRoot := math.Sqrt(a)
	fmt.Println("squareRoot = ", squareRoot)

	sine := math.Sin(a)
	fmt.Println("sine = ", sine)

	cosine := math.Cos(a)
	fmt.Println("cosine = ", cosine)

	tangent := math.Tan(a)
	fmt.Println("tangent = ", tangent)

	logarithm := math.Log(a)
	fmt.Println("log = ", logarithm)

	exponential := math.Exp(a)
	fmt.Println("exponential = ", exponential)

	absolute := math.Abs(a)
	fmt.Println("absolute = ", absolute)

	power := math.Pow(a, 2)
	fmt.Println("power = ", power)
}