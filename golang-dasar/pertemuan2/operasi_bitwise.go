package main

import (
	"fmt"
)

func main()  {
	a := 15
	b := 3

	// bitwise AND
	fmt.Println("a & b = ", a&b)

	// bitwise OR
	fmt.Println("a | b = ", a|b)

	// bitwise XOR
	fmt.Println("a ^ b = ", a^b)

	// bitwise negasi
	fmt.Println("^a = ", ^a)

	// bitwise shift left
	fmt.Println("a << 2 = ", a<<2)

	// bitwise shift right
	fmt.Println("a >> 2 = ", a>>2)
}