package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Selamat pagi!")
	case t.Hour() < 17:
		fmt.Println("Selamat sore")
	default:
		fmt.Println("Selamat malam")
	}
}

/*
Perintah "switch" tanpa kondisi
Perintah switch tanpa sebuah kondisi sama seperti switch true.

Konstruksi ini merupakan cara yang bersih untuk menulis rantaian if-then-else yang panjang.
*/
