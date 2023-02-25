package main

import "fmt"

const (
	// Buat bilangan yang besar dengan men-shift 1 bit ke kiri 100 kali.
	// Dengan kata lain, bilangan binari 1 diikuti dengan 100 angka nol.
	Big = 1 << 100
	// Shift kembali ke kanan sebanyak 99 kali, sehingga akhirnya menjadi
	// 1<<1, atau 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*
Konstanta Numerik
Konstanta numerik adalah nilai dengan presisi-tinggi.

Konstanta yang tidak bertipe menggunakan tipe yang dibutuhkan sesuai dengan konteksnya.

Coba cetak juga needInt(Big).

(Sebuah int dapat menyimpan sebanyak 64-bit integer, dan terkadang lebih sedikit, bergantung kepada mesin anda.)
*/
