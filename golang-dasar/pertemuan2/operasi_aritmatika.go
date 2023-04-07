package main

import "fmt"

func main()  {
	a := 15
	b := 5

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println("sum = ", sum)
	fmt.Println("sub = ", sub)
	fmt.Println("mul = ", mul)
	fmt.Println("div = ", div)
	fmt.Println("mod = ", mod)
}