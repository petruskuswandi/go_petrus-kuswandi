package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v tidak dapat digunakan disini
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
Kondisi "if" dan "else"
Variabel yang dideklarasikan dalam perintah singkat if juga dapat digunakan dalam blok else.

(Pada contoh, kedua pemanggilan terhadap fungsi pow dieksekusi sebelum main memanggil fmt.Println.)
*/
