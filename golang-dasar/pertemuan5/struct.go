package main

import "fmt"

type Orang struct {
	Name string
	Age int
}

func main()  {
	var orang1 Orang
	orang1.Name = "Tsaqib"
	orang1.Age = 20

	fmt.Println(orang1.Name)
	fmt.Println(orang1.Age)
}