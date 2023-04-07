package main

import "fmt"

func main()  {
	var a int = 10
	var ptr *int

	ptr = &a

	fmt.Println("Nilai a: ", a)
	fmt.Println("Alamat a: ", &a)
	fmt.Println("Nilai ptr: ", ptr)
	fmt.Println("Nilai yang ditunjuk ptr: ", *ptr)
}