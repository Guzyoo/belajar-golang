package main

import "fmt"

func main () {
	hasilTambah := tambah(80, 10)
	fmt.Println(hasilTambah)

	hasilKurang := kurang(40, 25)
	fmt.Println(hasilKurang)
}

func tambah(a int, b int) int {
	return a + b
}

func kurang(a int, b int) int {
	return a - b
}