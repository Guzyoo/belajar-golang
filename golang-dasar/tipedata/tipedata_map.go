package main

import "fmt"

func main()  {
	orang := map[string]string{
		"nama" : "Tsaqib",
		"alamat" : "Majalengka, Jawa Barat",
	}

	// len
	fmt.Println("len = ", len(orang))

	fmt.Println(orang)

	// merubah data
	orang["alamat"] = "Bandung, Jawa Barat"

	// delete
	delete(orang, "alamat")
	
	fmt.Println(orang)
}