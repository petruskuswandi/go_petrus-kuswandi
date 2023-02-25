package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))
}

/* Paket ("package")
Setiap program Go terbuat dari paket-paket (package).

Program mulai berjalan dalam paket main.

Program di sebelah menggunakan paket lain dengan meng- import "fmt" dan "math/rand".

Aturannya, nama paket sama dengan elemen terakhir dari path import. Sebagai contohnya, paket "math/rand" terdiri dari berkas-berkas yang dimulai dengan perintah package rand.

Catatan: Lingkungan di mana program di situs ini berjalan adalah deterministik, sehingga setiap kali anda mengeksekusi contoh program rand.Intn akan selalu mengembalikan angka yang sama. (Untuk mendapatkan angka yang acak, tambahkan pembangkit angka; lihat fungsi rand.Seed. Waktu adalah konstan di dalam playground, jadi anda perlu menggunakan hal lain sebagai pembangkit acak.)
*/
