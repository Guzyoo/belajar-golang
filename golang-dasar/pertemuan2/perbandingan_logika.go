package main

import "fmt"

func main()  {
	a := true
	b := false

	if a && b {
		fmt.Println("a dan b bernilai true")
	} else {
		fmt.Println("a dan b tidak bernilai true")
	}

	if a || b {
		fmt.Println("a atau b bernilai true")
	} else {
		fmt.Println("a atau b tidak bernilai true")
	}
}