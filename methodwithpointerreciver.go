package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	//If a method has a pointer receiver, then you don’t necessarily need to use the pointer dereferencing syntax (*e)
	//to get the value of the receiver. You can use simple e which will be the address of the value that pointer points
	//to but Go will understand that you are trying to perform an operation on the value itself and under the hood,
	//it will convert e to (*e).
	e.name = newName
}

func main() {
	e := Employee{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// change name
	e.changeName("Monica Geller")
	// e after name change
	fmt.Println("e after name change =", e)
}
