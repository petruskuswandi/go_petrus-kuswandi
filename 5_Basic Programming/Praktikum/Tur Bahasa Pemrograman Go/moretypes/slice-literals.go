package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/*
Inisialisasi Slice
Menginisialisasi slice mirip dengan array tapi tanpa mendefinisikan panjangnya.

Berikut ini adalah sebuah array:

[3]bool{true, true, false}
Dan berikut ini membuat array yang sama seperti di atas, kemudian membuat sebuah slice yang mengacu kepadanya:

[]bool{true, true, false}
*/
