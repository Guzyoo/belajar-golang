package main

import "fmt"

func filter(numbers []int, f func(int) bool) [] int  {
	var result []int
	for _, v := range numbers {
		if f (v) {
			result = append(result, v)
		}
	}
	return result
} 

func main()  {
	numbers := []int {3, 4, 5, 1, 7, 9, 10, 2, 6, 8}

	even := filter(numbers, func(x int) bool {
		return x % 2 == 0
	})

	fmt.Println(even)
}