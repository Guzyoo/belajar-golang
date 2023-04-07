package main

import "fmt"

func main()  {
	var num = 8

	if num < 0 {
		fmt.Println("num adalah bilangan negatif")
	} else if num > 0{
		fmt.Println("num adalah bilangan positif")
	} else {
		fmt.Println("num adalah bilangan nol")
	}
}