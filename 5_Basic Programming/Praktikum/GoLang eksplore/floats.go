package main

import (
	"fmt"
)

func main() {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)

	fmt.Print("\n")
	var a float64 = 1.7e+308
	fmt.Printf("Type: %T, value %v", a, a)
}
