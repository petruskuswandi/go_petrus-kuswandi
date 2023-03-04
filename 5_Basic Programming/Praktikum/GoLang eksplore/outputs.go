package main

import "fmt"

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	fmt.Print("\n", "\n")
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print("\n", "\n")
	fmt.Print(i, "\n", j)
	fmt.Print("\n")

	fmt.Print(i, " ", j)
	fmt.Print("\n")

	var k, l = 10, 20

	fmt.Print(k, l)
	fmt.Print("\n")
	fmt.Println(i, j)

	var m string = "Hello"
	var n int = 15

	fmt.Printf("m has value: %v and type: %T\n", m, m)
	fmt.Printf("n has value: %v and type: %T", n, n)
}
