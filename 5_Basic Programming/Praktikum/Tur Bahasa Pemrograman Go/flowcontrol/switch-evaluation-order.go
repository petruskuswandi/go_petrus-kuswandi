package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Kapan hari Sabtu?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Sekarang.")
	case today + 1:
		fmt.Println("Besok")
	case today + 2:
		fmt.Println("Dua hari lagi")
	default:
		fmt.Println("Masih jauh.")
	}
}

/*
Urutan evaluasi "switch"
Kondisi pada switch dievaluasi dari atas ke bawah, berhenti saat sebuah kondisi sukses.

Sebagai contoh,

switch i {
case 0:
case f():
}
tidak akan memanggil fungsi f jika i==0.

Catatan: Waktu dalam Go playground selalu berawal dari 2009-11-10 23:00:00 UTC, sebuah nilai yang maknanya bisa dicari oleh pembaca.
*/
