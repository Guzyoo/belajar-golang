package main

import "fmt"

func calculate(x, y int, operation func(int, int) int ) int {
	result := operation(x, y)
	return result
}

func main()  {
	hasilTambah := calculate(20, 5, func(x, y int) int {
		return x + y
	})
	hasilKali := calculate(20, 5, func (x, y int) int {
		return x * y
	})
	fmt.Println("Hasil Tambah", hasilTambah)
	fmt.Println("Hasil Kali", hasilKali)
}