package main

import "fmt"

type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}
type Material interface {
	Shape
	Object
}
type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return (c.side * c.side) * 6
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
func main() {
	c := Cube{3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Println("c", c)
	fmt.Println("m", m.(Cube))
	fmt.Println("s", s.(Cube))
	fmt.Println("o", o.(Cube))
	explain(c)
	explain(m)
	explain(s)
	explain(o)

}

func explain(i interface{}) {
	//Note : It has inherited multiple interfaces , so always will go in first case
	switch i.(type) {
	case Material:
		fmt.Println("m", i.(Material))
	case Object:
		fmt.Println("o", i.(Object))
	case Shape:
		fmt.Println("s", i.(Shape))


	default:
		fmt.Println("unknown", i)
	}
}
