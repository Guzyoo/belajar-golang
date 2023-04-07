package main

import "fmt"

type Orang struct {
	nama string
	umur int
}
func (o Orang) sayHalo() {
	fmt.Println("Haloo, nama Saya ",o.nama)
	fmt.Println("Usia Saya ",o.umur)
}

func main()  {
	o := Orang {nama: "Tsaqib", umur: 20}
	o.sayHalo()
}