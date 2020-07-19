package main

import "fmt"

func main() {
	var s []int
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = a[2:4]

	/*a[2] = 33
	a[3] = 44*/

	fmt.Println("before", s)
	fmt.Println("before", a)
	s[1] = 11
	s1 := append(s, 22, 33, 44, 55, 66, 77, 88, 99)
	fmt.Println("after", s)
	fmt.Println("after S1", s1)
	fmt.Println("after", a)
}
