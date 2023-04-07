package main

import (
	"fmt"
	"math"
)

func main()  {
	// Menghitung absolut -10
	fmt.Println("Nilai Absolut dari - 10 = ", math.Abs(-10))

	// Membulatkan angka 3.14 ke atas
	fmt.Println("Nilai 3.14 dibulatkan ke atas = ", math.Ceil(3.14))

	// Menghitung nilai cosinus dari 45 derajat
	fmt.Println("Nilai cosinus dari 45 derajat = ", math.Cos(math.Pi/4))

	// Menghitung eksponensial dari 2
	fmt.Println("Nilai eksponen dari 2 = ", math.Exp(2))

	// Membulatkan 3.14 ke bawah
	fmt.Println("Nilai 3.14 dibulatkan ke bawah = ", math.Floor(3.14))

	// Menghitung Nilai logaritma natural dari 10 dengan basis 2
	fmt.Println("Nilai logaritma natural dari 10 dengan basis 2 = ", math.Log2(10))

	// Mengembalikan nilai maksimum dari dua angka
	fmt.Println("Nilai maksimum dari 10 dan 20 = ", math.Max(10, 20))

	// Mengembalikan nilai minimum dari dua angka
	fmt.Println("Nilai minimum dari 10 dan 20 = ", math.Min(10, 20))

	// Menghitung nilai pangkat dari 2^3
	fmt.Println("Nilai pangkat dari 2^3 = ", math.Pow(2, 3))

	// Menghitung sinus dari 30 derajat
	fmt.Println("Nilai sinus dari 30 derajat = ", math.Sin(math.Pi/6))

	// Menghitung nilai akar kuadrat dari 16
	fmt.Println("Nilai akar kuadrat dari 16 = ", math.Sqrt(16))

	// Menghitung nilai tangen dari 60 derajat
	fmt.Println("Nilai tangen dari 60 derajat = ", math.Tan(math.Pi/3))
}