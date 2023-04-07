package main

import (
	"fmt"
	"os"
)

func main()  {
	hostname, _ := os.Hostname()
	fmt.Println("Host Name = ", hostname)

	username := os.Getenv("USERNAME")
	fmt.Println("Username = ", username)

	wd, _ := os.Getwd()
	fmt.Println("Working Directory = ", wd) 
}