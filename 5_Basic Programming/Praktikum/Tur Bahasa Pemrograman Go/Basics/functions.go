package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*
Fungsi
Sebuah fungsi bisa tanpa argumen atau banyak argumen.

Pada contoh ini, add memiliki dua parameter dengan tipe int.

Perhatikan bahwa tipe berada setelah nama variabel.

(Untuk mengetahui kenapa tipe ditulis seperti itu, lihat artikel Go tentang deklarasi sintaks.)
*/
