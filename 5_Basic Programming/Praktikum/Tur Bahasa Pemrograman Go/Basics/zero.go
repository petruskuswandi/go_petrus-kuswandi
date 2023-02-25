package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
Nilai kosong
Variabel yang dideklarasikan tanpa nilai awal diberikan nilai kosong.

Nilai kosong adalah:

0 untuk tipe numerik,
false untuk tipe boolean, dan
"" (string kosong) untuk string.
*/
