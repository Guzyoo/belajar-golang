package main

import "fmt"

func main()  {
	x := 5
	y := 15
	z := 35

	// AND (&&)
	if x < y && y < z {
		fmt.Println("kondisi terpenuhi")
	} else {
		fmt.Println("kondisi tidak terpenuhi")
	}

	// OR (||)
	if x > y || y < z {
		fmt.Println("Kondisi terpenuhi")
	} else {
		fmt.Println("Kondisi tidak terpenuhi")
	}

	// NOT (!)
	if ! (x > y) {
		fmt.Println("Kondisi terpenuhi")
	} else {
		fmt.Println("Kondisi tidak terpenuhi")
	}
}