package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*
Fungsi (lanjutan)
Saat dua atau lebih parameter fungsi memiliki tipe yang sama, anda bisa menggabungkannya menjadi satu.

Pada contoh ini, kita singkat

x int, y int
menjadi

x, y int
*/
