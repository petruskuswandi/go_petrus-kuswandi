package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

/*
Array
Deklarasi tipe dengan [n]T adalah untuk array dengan jumlah n dan bertipe T.

Ekspresi

var a [10]int
mendeklarasikan sebuah variabel a sebagai sebuah array dari sepuluh integer.

Panjang sebuah array adalah bagian dari tipenya, jadi array tidak bisa diubah ukurannya. Hal ini sepertinya membatasi, tapi jangan khawatir; Go menyediakan cara yang mudah untuk bekerja dengan array.
*/
