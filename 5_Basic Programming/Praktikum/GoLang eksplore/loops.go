package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// The continue Statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// The break Statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// Nested loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	fmt.Printf("\n")

	//The Range Keyword
	fruits = [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	fmt.Printf("\n")

	fruits = [3]string{"apple", "orange", "banan"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	fmt.Printf("\n")

	fruits = [3]string{"apple", "orange", "banan"}
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
	fmt.Printf("\n")

	char := "Hello"

	for i := 0; i < len(char); i++ {
		fmt.Printf(string(char[i]) + "-")
	}

	fmt.Println("	------>>>")
	for pos, char := range char {
		fmt.Printf("Character %c start at byte position %d\n", char, pos)
	}
}
