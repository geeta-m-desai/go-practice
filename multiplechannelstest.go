package main

import (
	. "fmt"
	. "time"
)

func square(c chan int) {
	Println("[square] reading")
	num := <-c
	Sleep(100000*1000)
	Println("Woke up after sleep")
	c <- num * num

}

func cube(c chan int) {
	Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	Println("[main] resuming")
	Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	Println("[main] resuming")
	Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	Println("squareVal",squareVal)
	Println("[main] sum of square and cube of", testNum, " is", sum)
	Println("[main] main() stopped")
}
