package main

import (
	"fmt"
	"math"
)

func main() {
	// Inisialisasi matriks
	var matrix [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Hitung jumlah diagonal utama dan kedua
	var mainSum, secondarySum int = masukanDiagonalSums(matrix)

	// Hitung selisih absolut antara jumlah diagonal
	var difference float64 = math.Abs(float64(mainSum - secondarySum))

	// Tampilkan hasil
	fmt.Printf("Selisih absolut antara jumlah diagonal matriks: %f\n", difference)
}

// Fungsi untuk menghitung jumlah diagonal utama dan diagonal kedua matriks
func masukanDiagonalSums(matrix [][]int) (int, int) {
	var mainSum, secondarySum int = 0, 0

	for i := 0; i < len(matrix); i++ {
		mainSum += matrix[i][i]
		secondarySum += matrix[i][len(matrix)-1-i]
	}

	return mainSum, secondarySum
}
