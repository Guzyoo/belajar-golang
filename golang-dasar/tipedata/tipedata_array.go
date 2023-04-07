package main

import "fmt"

func main()  {
	// bisa string atau int
	var a [3] string
	a [0] = "Haloo" 
	a [1] = "Selamat" 
	a [2] = "Datang !" 

	fmt.Println("a : ", a)
	fmt.Println("a[0] : ", a[0])
	fmt.Println("a[1] : ", a[1])
	fmt.Println("a[2] : ", a[2])

	// int
	var b [3]int
	b [0] = 1
	b [1] = 180
	b [2] = 90

	fmt.Println("b : ", b)
	fmt.Println("b [0] : ", b[0])
	fmt.Println("b [1] : ", b[1])
	fmt.Println("b [2] : ", b[2])

	// float
	var c [3]float64
	c [0] = 8.6
	c [1] = 7.5
	c [2] = 6.7

	fmt.Println("c : ", c)
	fmt.Println("c [0] : ", c[0])
	fmt.Println("c [1] : ", c[1])
	fmt.Println("c [2] : ", c[2])

	// bool
	var d [2]bool
	d [0] = true
	d [1] = false

	fmt.Println("d : ", d)
	fmt.Println("d [0] : ", d[0])
	fmt.Println("d [1] : ", d[1])

	// len
	fmt.Println("Panjang a = ", len(a))
	fmt.Println("Panjang b = ", len(b))
	fmt.Println("Panjang c = ", len(c))
	fmt.Println("Panjang d = ", len(d))
}