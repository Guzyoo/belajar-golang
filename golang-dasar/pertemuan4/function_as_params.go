package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func calculate(x, y int, operation func (int, int) int ) int {
	result := operation(x, y)
	return result
}

func main()  {
	x, y := 6, 12
	
	fmt.Println("Hasil Penjumlahan = ", calculate(x, y, add))
	fmt.Println("Hasil Perkalian = ", calculate(x, y, multiply))
}