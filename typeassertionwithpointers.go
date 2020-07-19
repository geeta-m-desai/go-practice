package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{5.0, 4.0}
	r1 := Rect{5.0, 4.0}
	r2 := Rect{5.0, 3.0}
	s := Shape(&r)
	newS := Shape(&r1)
	altS := Shape(&r)

	area := s.Area()
	fmt.Println("area of rectangle is", area)
	//Seems like Go is happy to pass a copy of the pointer value as the receiver to the Perimeter method perhaps because
	//it is not a very dangerous idea, it is just a copy and nobody can mutate it accidentally.
	perimeter := s.Perimeter()
	fmt.Println("perimeter of rectangle is", perimeter)
	//Comparision of interfaces
	fmt.Println(s == newS)
	fmt.Println(s == altS)
	fmt.Println(r == r1)
	fmt.Println(r == r2)

}
