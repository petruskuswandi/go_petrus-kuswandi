package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))
}

/*
Kondisi "if" singkat
Seperti pada perintah for, perintah if bisa diawali dengan perintah singkat untuk dieksekusi sebelum kondisi.

Variabel yang dideklarasikan pada perintah singkat tersebut hanya berlaku sampai lingkup kondisi if berakhir.

(Coba gunakan v di akhir perintah return, dan jalankan)
*/
