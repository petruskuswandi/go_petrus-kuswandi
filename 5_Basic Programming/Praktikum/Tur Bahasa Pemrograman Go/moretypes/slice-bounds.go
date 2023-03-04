package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

/*
Nilai Default Slice
Saat memotong, anda bisa mengindahkan batas bawah atau atas sehingga Go akan menggunakan nilai default-nya.

Nilai default-nya adalah nol untuk batas bawah dan panjang dari slice untuk batas atas.

Untuk sebuah array berikut,

var a [10]int
ekspresi slice berikut adalah setara:

a[0:10]
a[:10]
a[0:]
a[:]
*/
