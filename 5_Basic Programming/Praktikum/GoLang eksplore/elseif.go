package main

import (
	"fmt"
)

func main() {
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}

	x := 30
	if x >= 10 {
		fmt.Println("x is larger than or equal to 10")
	} else if x > 20 {
		fmt.Println("x is larger than 20.")
	} else {
		fmt.Println("x is less than 10.")
	}
}
