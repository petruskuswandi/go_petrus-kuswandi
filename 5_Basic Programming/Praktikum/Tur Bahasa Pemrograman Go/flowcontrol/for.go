package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
Pengulangan ("for")
Go hanya memiliki satu perintah pengulangan yaitu for.

Dasar dari pengulangan for yaitu memiliki tiga komponen yang dipisahkan oleh titik-koma:

perintah awal: dieksekusi sebelum iterasi pertama
ekspresi kondisi: dievaluasi sebelum iterasi pertama
perintah akhir: dieksekusi disetiap akhir iterasi
Perintah awal biasanya berupa deklarasi variabel singkat, dan variabel yang dideklarasikan tersebut hanya dapat digunakan dalam skop perintah for.

Pengulangan akan berhenti saat ekspresi kondisi bernilai false.

Catatan: Tidak seperti bahasa C, Java, atau JavaScript, tidak ada tanda kurung yang digunakan menutupi ketiga komponen dari perintah for dan tanda kurung kurawal { } selalu dibutuhkan.
*/
