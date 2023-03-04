package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
	return x + y
}

func myFunction1(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunction3(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction4(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " world!"
	return
}

func myFunction5(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " world!"
	return
}

func myFunction6(x int, y string) (result1 int, txt2 string) {
	result1 = x + x
	txt2 = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction(1, 2))

	fmt.Println(myFunction1(2, 2))

	fmt.Println(myFunction2(2, 3))

	total := myFunction3(14, 5)
	fmt.Println(total)

	fmt.Println(myFunction4(5, "Hello"))

	a, b := myFunction5(5, "Hello")
	fmt.Println(a, b)

	_, d := myFunction6(5, "Hello")
	fmt.Println(d)

	c, _ := myFunction6(5, "Hello")
	fmt.Println(c)
}
