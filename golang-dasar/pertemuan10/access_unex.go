package main

import "fmt"

// Unexported Variable
var unexportedVariable string = "Ini adalah Unexported Variable"

// Exported Variable
var ExportedVariable string = "Ini adalah Exported Variable"

func main(){
	fmt.Println(unexportedVariable)
	fmt.Println(ExportedVariable)
}