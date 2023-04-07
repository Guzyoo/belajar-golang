package main

import "fmt"

func main()  {
	arr := []string{"satu", "dua", "tiga"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}