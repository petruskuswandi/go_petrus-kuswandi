package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go berjalan pada")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		// plan9, windwos...
		fmt.Printf("%s. \n", os)
	}
}

/*
Perintah "switch"
Perintah switch untuk mempermudah membuat beberapa perintah kondisi if - else. Go akan menjalankan case pertama yang nilainya sama dengan ekspresi kondisi yang diberikan.

Perintah switch pada Go hampir sama dengan bahasa C, C++, Java, Javascript, dan PHP; hanya saja pada Go akan menjalankan case yang terpilih, bukan semua case yang ada selanjutnya. Efeknya, perintah break yang biasanya dibutuhkan diakhir setiap case pada bahasa lainnya dibuat secara otomatis oleh Go. Perbedaan penting lainnya yaitu ekspresi kondisi case pada Go tidak harus konstanta, dan nilainya tidak harus integer.
*/
