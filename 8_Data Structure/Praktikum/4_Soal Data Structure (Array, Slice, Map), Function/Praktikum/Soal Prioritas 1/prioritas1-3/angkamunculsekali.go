package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	// Buat kamus kosong untuk menyimpan frekuensi masing-masing angka
	frekuensi := make(map[rune]int)

	// Hitung frekuensi masing-masing angka dalam input
	for _, num := range angka {
		frekuensi[num]++
	}

	// Buat slice kosong untuk menyimpan angka yang hanya muncul 1 kali
	uniqueNums := []int{}

	// Cek setiap angka dalam kamus, jika frekuensinya 1, maka angka tersebut hanya muncul 1 kali
	for num, count := range frekuensi {
		if count == 1 {
			intNum, _ := strconv.Atoi(string(num))
			uniqueNums = append(uniqueNums, intNum)
		}
	}

	// Kembalikan slice angka yang hanya muncul 1 kali
	return uniqueNums
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
