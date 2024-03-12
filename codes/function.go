package main

import (
	"fmt"
)

func beyondHello() {
	var x int
	x = 3

	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}
