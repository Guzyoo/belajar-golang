package main

import (
	"fmt"
	"sort"
)

func main()  {
	numbers := []int{4, 7, 9, 2, 1}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers [i] < numbers [j]
	})

	fmt.Println(numbers)
}