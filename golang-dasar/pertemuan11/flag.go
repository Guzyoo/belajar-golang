package main

import (
	"flag"
	"fmt"
)

func main()  {
	var (
		name string
		age int
		isFemale bool
	)

	flag.StringVar(&name, "name", "Rissya", "a name to say hello to")
	flag.IntVar(&age, "age", 0, "your age")
	flag.BoolVar(&isFemale, "Female", false, "a gender")
	flag.Parse()

	fmt.Println("Hello", name)
	fmt.Println("Your age is:", age)
	fmt.Println("Your gender:", isFemale)
}