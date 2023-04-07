package main

import "fmt"

func main()  {
	var angka int = 4

	switch angka {
	case 1:
		fmt.Println("Hari ke - ", angka, "adalah Ahad")
	case 2:
		fmt.Println("Hari ke - ", angka, "adalah Senin")
	case 3:
		fmt.Println("Hari ke - ", angka, "adalah Selasa")
	case 4:
		fmt.Println("Hari ke - ", angka, "adalah Rabu")
	case 5:
		fmt.Println("Hari ke - ", angka, "adalah Kamis")
	case 6:
		fmt.Println("Hari ke - ", angka, "adalah Jum'at")
	case 7:
		fmt.Println("Hari ke - ", angka, "adalah Sabtu")
	default:
		fmt.Println("Angka yang Anda masukkan salah atau tidak valid")
	}
}