package main

import (
	"fmt"
	"regexp"
)

func main ()  {
	// mencocokkan pola pada sebuah string
	matched, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(matched)

	// mencocokkan sebuah pola pada sebuah string dan mengembalikan submatch
	r :=  regexp.MustCompile("a([a-z]+)e")
	submatched := r.FindStringSubmatch("apple")
	fmt.Println(submatched)
}