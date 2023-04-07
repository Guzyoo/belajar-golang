package main

import "fmt"

func main()  {
	// named
	type Person string
	var nama Person = "Tsaqib"
	fmt.Println(nama)

	// alias
	type PersonAlias = string
	var namaAlias PersonAlias = "Ilham"
	fmt.Println(namaAlias)

	// struct
	type PersonStruct struct {
		nama string
		umur int
	}

	var orang PersonStruct = PersonStruct{
		nama: "Ilham",
		umur: 20,
	}

	fmt.Println(orang)
}