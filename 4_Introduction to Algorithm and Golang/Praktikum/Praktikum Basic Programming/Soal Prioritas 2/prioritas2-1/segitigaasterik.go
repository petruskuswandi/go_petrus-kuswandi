package main

import (
	"fmt"
)

func main() {
	var baris int
	fmt.Print("Masukkan jumlah baris segitiga asterik: ")
	fmt.Scan(&baris)

	for i := 1; i <= baris; i++ {
		for j := i; j <= baris; j++ {
			fmt.Print(" ")
		}
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
