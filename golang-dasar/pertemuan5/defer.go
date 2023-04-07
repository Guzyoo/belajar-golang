package main

import "fmt"

func main()  {
	defer fmt.Println("Tulisan ini muncul di akhir")
	fmt.Println("Ini muncul pertama")
}