package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) ChangeName (name string) {
	p.Name = name
}

func main()  {
	person := Person{Name: "Tsaqib", Age: 20}
	fmt.Println(person.Name)
	person.ChangeName("Tsaqib Ilham Nur")
	fmt.Println(person.Name)
}