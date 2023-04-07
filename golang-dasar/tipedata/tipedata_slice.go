package main

import "fmt"

func main() {
	hari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Ahad"}

	slice := hari[1:4]

	fmt.Println(hari)
	fmt.Println(slice)
}