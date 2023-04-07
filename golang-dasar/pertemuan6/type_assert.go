package main

import "fmt"

func main()  {
	list := []interface{}{"Halo", 44, true}

	for _, val := range list {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Print(val.(int))
		default:
			fmt.Println("Unknown type")
		}
	}
}