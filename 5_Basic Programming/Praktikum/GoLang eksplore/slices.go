package main

import (
	"fmt"
)

func main() {
	/*
		Create a Slice With []datatype{values}

		Syntax
		slice_name := []datatype{values}

		myslice := []int{}
	*/

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	fmt.Printf("\n")

	//Create a Slice From an Array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	fmt.Printf("\n")

	//Create a Slice With The make() Function
	//slice_name := make([]type, length, capacity)
	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	myslice4 := make([]int, 5)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
}
