package main

import "fmt"

type person struct {
	name string
	dept string
}

func main() {
	var count int
	var name string
	var dept string

	fmt.Print("Enter the total no. of person: ")
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter the name of person %d : ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Enter the department of person %d : ", i+1)
		fmt.Scanln(&dept)
		s := person{name: name, dept: dept}
		fmt.Printf("\n person %d: %s ", i+1, s)
		fmt.Printf("\n\n")
	}
}
