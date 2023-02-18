package main

import (
	"fmt"
)

func main() {
	var bawah, tinggi, atas int

	fmt.Print("Masukkan nilai atas: ")
	fmt.Scan(&atas)
	fmt.Print("Masukkan nilai bawah: ")
	fmt.Scan(&bawah)
	fmt.Print("Masukkan nilai tinggi: ")
	fmt.Scan(&tinggi)

	luas := (atas + bawah) / 2 * tinggi

	fmt.Print("Luas trapesium = ", luas)
}
