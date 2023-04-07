package main

import (
	"errors"
	"fmt"
)

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0")
	}
	return x / y, nil
}

func main()  {
	result, err := division(20, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}