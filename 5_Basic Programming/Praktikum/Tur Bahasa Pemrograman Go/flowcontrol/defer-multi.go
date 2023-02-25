package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/*
Penundaan bertumpuk
Fungsi yang dipanggil dengan defer di- push ke sebuah stack. Saat fungsi berakhir, panggilan yang tadi ditunda dieksekusi dengan urutan last-in-first-out (yang terakhir masuk menjadi pertama keluar).

Untuk belajar lebih lanjut tentang perintah defer bacalah blog berikut.
*/
