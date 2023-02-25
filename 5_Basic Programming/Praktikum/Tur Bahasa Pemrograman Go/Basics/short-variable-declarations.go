package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

/*
Deklarasi variabel singkat
Dalam sebuah fungsi, perintah singkat := bisa digunakan menggantikan var dengan tipe implisit.

Di luar sebuah fungsi, setiap perintah harus dimulai dengan kata kunci (var, func, dst), sehingga := tidak bisa digunakan.
*/
