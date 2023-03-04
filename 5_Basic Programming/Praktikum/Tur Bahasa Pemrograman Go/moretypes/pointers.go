package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // menunjuk ke i
	fmt.Println(*p) // baca i lewat pointer
	*p = 21         // set i lewat pointer
	fmt.Println(i)  // lihat nilai terbaru dari i

	p = &j         // p menunjuk ke j
	*p = *p / 37   // bagi nilai j lewat pointer
	fmt.Println(j) // lihat nilai terbaru dari j
}

/*
Pointer
Go memiliki pointer. Sebuah pointer menyimpan alamat dari sebuah nilai.

Tipe *T adalah pointer ke sebuah nilai T. Nilai kosong dari pointer adalah nil.

var p *int
Operator & mengembalikan operan pointer dari variabel.

i := 42
p = &i
Operator * mengembalikan nilai yang ditunjuk oleh pointer.

fmt.Println(*p) // baca nilai i lewat pointer p
*p = 21         // set nilai i lewat pointer p
Cara ini disebut dengan "dereferencing" atau "indirecting".

Tidak seperti C, Go tidak memiliki fungsi aritmatika pada pointer.
*/
