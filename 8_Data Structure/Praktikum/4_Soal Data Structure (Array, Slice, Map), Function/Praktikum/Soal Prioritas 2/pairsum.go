package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// Buat kamus kosong untuk menyimpan indeks masing-masing angka dalam array
	idxMap := make(map[int]int)

	// Loop melalui setiap angka dalam array
	for i, num := range arr {
		// Hitung selisih antara target dan angka saat ini
		diff := target - num

		// Cek apakah selisih sudah ada dalam kamus, jika ada maka kembalikan indeks saat ini dan indeks selisih tersebut
		if idx, ok := idxMap[diff]; ok {
			return []int{idx, i}
		}

		// Tambahkan indeks angka saat ini ke kamus
		idxMap[num] = i
	}

	// Jika tidak ada pasangan yang ditemukan, kembalikan slice kosong
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
