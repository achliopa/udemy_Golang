package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// 3rd way of declaring structs
	// var alex person //declaration and zero value assignment
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// alex.firstName = "Alex"
	// alex.lastname = "Anderson"
	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
