package main

import (
	"fmt"
)

func main() {
	/*
		Declare an Array
		1. With the var keyword
		Syntax
		var array_name = [length]datatypes{values} // here length is defined

		or

		var array_name = [...]datatypes{values} //here length is inferred

		2. With the := sign
		array_name := [length]datatypes{values} // here length is defined

		or

		array_name := [...]datatypes{values} // here length is inferred
	*/

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Printf("\n")

	var arr3 = [...]int{9, 10, 11}
	arr4 := [...]int{12, 13, 14, 15, 16}

	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Printf("\n")

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)

	fmt.Printf("\n")
	fmt.Printf("\n")

	// Access Elements of an Array
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	fmt.Printf("\n")

	prices[2] = 50
	fmt.Println(prices)

	fmt.Printf("\n")

	arr5 := [5]int{}              // not initialized
	arr6 := [5]int{1, 2}          // partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	fmt.Printf("\n")

	arr8 := [5]int{1: 10, 2: 40}

	fmt.Println(arr8)

	fmt.Printf("\n")

	arr9 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr10 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr9))
	fmt.Println(len(arr10))
}
