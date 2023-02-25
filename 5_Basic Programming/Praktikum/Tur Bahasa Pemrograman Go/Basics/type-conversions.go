package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/*
Konversi Tipe
Ekspresi T(v) mengkonversi nilai v ke tipe T.

Beberapa konversi numerik:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Atau, sederhananya:

i := 42
f := float64(i)
u := uint(f)
Tidak seperti C, perintah pengisian nilai pada Go antara item dengan tipe berbeda membutuhkan konversi eksplisit. Coba hilangkan konversi float64 atau int pada contoh sebelah dan lihat apa yang terjadi.
*/
