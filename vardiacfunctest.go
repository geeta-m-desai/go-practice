package main

import "fmt"

func getMultiples(factor int, args ...int) []int {
	for index, val := range args {
		args[index] = val * factor
	}

	return args
}

func main() {
	s := []int{10, 20, 30}

	mult1 := getMultiples(2, s...)

	fmt.Println(s, mult1)
}
