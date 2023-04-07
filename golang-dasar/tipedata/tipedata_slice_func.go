package main

import "fmt"

func main()  {
	hari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Ahad"}

	slice := hari[1:]

	// slice[0] = "Selasa Baru"
	// slice[1] = "Rabu Lagi"

	fmt.Println(hari)
	fmt.Println(slice)

	// append
	slice2 := append(slice, "Sabtu Baru")
	fmt.Println(slice2)

	// make
	slice3 := make([]string, 2, 5)
	slice3[0] = "Senin Baru"
	slice3[1] = "Selasa Baru"
	fmt.Println(slice3)

	// copy
	copy(slice, slice2)
	fmt.Println(slice)
}