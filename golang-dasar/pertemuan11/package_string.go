package main

import (
	"fmt"
	"strings"
)

func main()  {
	// Contains
	fmt.Println(strings.Contains("Hello Dunia", "Dunia"))
	fmt.Println(strings.Contains("Hello Dunia", "dunia"))

	// Count
	fmt.Println(strings.Count("Dunia ini indah", "i"))

	// Replace
	fmt.Println(strings.Replace("Hello Dunia", "Dunia", "Indonesia", 1))
	fmt.Println(strings.Replace("Hello Dunia", "a", "i", - 1))

	// ToUpper dan ToLower
	fmt.Println(strings.ToUpper("Hello Dunia"))
	fmt.Println(strings.ToLower("Hello Dunia"))

	// TrimSpace
	fmt.Println(strings.TrimSpace(" Hello Dunia "))
}