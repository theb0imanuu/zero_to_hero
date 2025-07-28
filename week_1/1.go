package main

import (
	"fmt"
	"math"
	"time"
)

// We have to start somewhere don't we😜
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello there")

	//Printing the time
	fmt.Println("The time is", time.Now())

	//Let's play around with the math package
	fmt.Println("The square root 49 is", math.Sqrt(49))

	fmt.Println(add(2, 3))
}
