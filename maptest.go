package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}
	minaAge, minaOk := age["mina"]
	jessyAge, jessyOk := age["jessy"]

	fmt.Println(minaAge, minaOk)
	fmt.Println(jessyAge, jessyOk)
}

