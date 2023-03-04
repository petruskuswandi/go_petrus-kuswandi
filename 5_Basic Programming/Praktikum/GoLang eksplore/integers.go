package main

import (
	"fmt"
)

func main() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value %v", y, y)

	var a uint = 500
	var b uint = 4500
	fmt.Printf("\n")
	fmt.Printf("Type: %T, value %v\n", a, a)
	fmt.Printf("Type: %T, value %v", b, b)
}
