package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main()  {
	fmt.Println(sum(1, 4, 5, 6, 2, 9, 10))
	fmt.Println(sum(90, 80, 12, 40, 70))
}