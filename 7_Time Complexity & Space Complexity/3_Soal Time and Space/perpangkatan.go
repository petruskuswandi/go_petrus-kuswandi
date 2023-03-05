package main

import (
	"fmt"
)

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n%2 == 0 {
		return pow(x*x, n/2)
	} else {
		return x * pow(x*x, (n-1)/2)
	}
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
	fmt.Println(pow(1, 10)) // 1
}
