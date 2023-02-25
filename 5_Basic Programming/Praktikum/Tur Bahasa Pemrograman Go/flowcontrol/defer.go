package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/*
Perintah "defer"
Perintah defer menunda eksekusi dari sebuah fungsi sampai fungsi yang melingkupinya selesai.

Argumen untuk pemanggilan defer dievaluasi langsung, tapi pemanggilan fungsi tidak dieksekusi sampai fungsi yang melingkupinya selesai.
*/
