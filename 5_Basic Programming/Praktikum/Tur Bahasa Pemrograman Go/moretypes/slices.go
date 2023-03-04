package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

/*
Potongan ("slice")
Sebuah array memiliki ukuran yang tetap. Sebuah slice ukurannya bisa dinamis, bisa mengacu secara fleksibel ke elemen dalam sebuah array. Dalam praktiknya, slice lebih sering digunakan daripada array.

Tipe []T adalah sebuah slice dengan elemen bertipe T.

Sebuah slice dibentuk dengan menspesifikasikan dua indeks, batas bawah dan batas atas, dipisahkan oleh sebuah tanda titik-dua:

a[bawah : atas]
Notasi di atas memotong rentang dari slice a yang mengikutkan elemen bawah, tapi tidak memasukan bagian terakhir (atas).

Ekspresi berikut membuat sebuah slice yang mengikutkan elemen 1 sampai 3 dari slice a:

a[1:4]
*/
