package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
Perintah for adalah "while"-nya Go
Dengan cara ini anda bisa menghilangkan perintah awal dan akhir, menggunakan hanya ekpresi kondisi sehingga for menjadi seperti while pada bahasa C.
*/
