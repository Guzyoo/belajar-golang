package main

import (
	"calculation"
	"fmt"
)

func main () {
	var sisi float64 = 5
	luas := calculation.SquareArea(sisi)
	keliling := calculation.SquareParameter(sisi) 

	fmt.Println("Luas Persegi = ", luas)
	fmt.Println("Keliling Persegi = ", keliling)
}