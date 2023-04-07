package main

import "fmt"

func main()  {
	x := 20
	y := 40

	if x == y {
		fmt.Println("x sama dengan y")
	} else {
		fmt.Println("x tidak sama dengan y")
	}

	if x < y {
		fmt.Println("x kurang dari y")
	} else {
		fmt.Println("x tidak kurang dari y")
	}
}