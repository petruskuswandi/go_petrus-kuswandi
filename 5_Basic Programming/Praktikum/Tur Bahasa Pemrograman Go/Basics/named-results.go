package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/*
Fungsi dengan kembalian bernama
Nilai kembalian dari fungsi bisa diberi nama. Jika nilai kembalian memiliki nama, mereka diperlakukan seperti variabel yang didefinisikan pada fungsi.

Penamaan tersebut seharusnya digunakan untuk mendokumentasikan makna dari nilai kembalian.

Sebuah perintah return tanpa nilai kembalian, mengembalikan nilai terakhir yang disimpan dari setiap variabel kembalian. Hal ini dikenal dengan kembalian "naked".

Perintah kembalian "naked" sebaiknya hanya digunakan pada fungsi yang singkat, seperti pada contoh di sebelah, karena bisa mengganggu pembacaan kode bila digunakan pada fungsi yang panjang.
*/
