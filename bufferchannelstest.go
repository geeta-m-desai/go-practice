package main

import (
	"fmt"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int,3)

	go squares(c)

	c <- 1
	fmt.Println("main() started 1")
	c <- 2
	fmt.Println("main() started 2")
	c <- 3
	fmt.Println("main() started 3")
	//Without overflow of buffer , go routine will not get chance to execute
	c <- 4 // blocks here

	fmt.Println("main() stopped")
}
