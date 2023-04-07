package main

import "fmt"

func main()  {
	var a int32 = 10000

	var b int64 = int64(a)
	var c int16 = int16(a)

	var d float64 = float64(a)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("b = ", c)
	fmt.Println("b = ", d)
}