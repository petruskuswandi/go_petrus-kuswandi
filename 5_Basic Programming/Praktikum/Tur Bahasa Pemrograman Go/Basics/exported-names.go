package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

/*
Nama-nama yang diekspor
Pada bahasa Go, sebuah nama diekspor jika diawali dengan huruf besar. Sebagai contohnya, Pizza itu nama yang diekspor, sebagaimana juga Pi, yang diekspor dari paket math.

pizza dan pi tidak diawali dengan huruf besar, maka mereka tidak diekspor.

Saat mengimpor sebuah paket, anda hanya bisa mengacu pada nama yang diekspor. Setiap nama yang tidak diekspor tidak akan bisa diakses dari luar paket.

Jalankan kode di sebelah. Perhatikan akan adanya pesan kesalahan.

Untuk memperbaikinya, ganti nama math.pi dengan math.Pi dan coba kembali.
*/
