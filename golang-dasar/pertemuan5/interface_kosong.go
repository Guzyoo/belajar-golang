package main

import "fmt"

func print(val interface{}) {
	fmt.Println(val)
} 

func main () {
	print("Haloo semuanya")
	print(591)
	print(true)
}