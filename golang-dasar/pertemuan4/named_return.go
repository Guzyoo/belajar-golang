package main

import "fmt"

func sumAndDifference (a int, b int) (sum int, difference int){
	sum = a + b
	difference = a - b
	return
}
func main()  {
	sum, difference := sumAndDifference(25, 50)
	fmt.Println("Sum = ", sum)
	fmt.Println("Difference = ", difference)
}