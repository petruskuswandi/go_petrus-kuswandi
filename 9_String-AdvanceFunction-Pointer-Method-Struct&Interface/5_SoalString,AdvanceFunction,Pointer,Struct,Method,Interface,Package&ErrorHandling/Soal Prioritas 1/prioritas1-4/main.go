package main

import "fmt"

func findMinMax(numbers *[6]int) (int, int) {
	var max, min int
	for i, num := range *numbers {
		if i == 0 || num > max {
			max = num
		}
		if i == 0 || num < min {
			min = num
		}
	}
	return max, min
}

func main() {
	var numbers [6]int
	for i := 0; i < 6; i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		fmt.Scanln(&numbers[i])
	}
	max, min := findMinMax(&numbers)
	fmt.Printf("Nilai maksimum: %d\n", max)
	fmt.Printf("Nilai minimum: %d\n", min)
}
