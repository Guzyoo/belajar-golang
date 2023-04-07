package main

import "fmt"

func main()  {
	var a int = 10
	var b int = 20

	fmt.Println("Sebelum swipe, nilai a:", a, "dan nilai b:", b)
	
	swap(&a, &b)

	fmt.Println("Setelah swipe, nilai a:", a, "dan nilai b:", b)
}

func swap (x *int, y *int) {
	var tempt int

	tempt = *x
	*x = *y
	*y = tempt
}