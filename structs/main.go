package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {

	alex := person{
		"Michal",
		"Baranowski",
		contactInfo{
			"bombell@gmail.com",
			16010,
		},
	}

	alex.updateName("Jimmy")
	alex.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}