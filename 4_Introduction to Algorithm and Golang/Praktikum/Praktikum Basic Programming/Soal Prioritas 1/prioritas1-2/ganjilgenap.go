package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}
