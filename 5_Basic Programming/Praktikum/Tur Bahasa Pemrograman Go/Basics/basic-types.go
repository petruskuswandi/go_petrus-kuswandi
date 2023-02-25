package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)
}

/*
Tipe dasar
Tipe dasar pada Go yaitu,

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias untuk uint8

rune // alias untuk int32
     // merepresentasikan sebuah kode Unicode

float32 float64

complex64 complex128
Contoh sebelah memperlihatkan variabel dengan beberapa tipe, dan juga deklarasi variabel bisa "difaktorkan" ke dalam blok, seperti pada perintah import.

Tipe int, uint, dan uintptr biasanya memiliki panjang 32 bits pada sistem 32-bit dan 64 bits pada sistem 64-bit. Saat anda membutuhkan nilai integer anda sebaiknya menggunakan tipe int kecuali anda memiliki alasan tertentu menggunakan tipe berukuran khusus atau unsigned integer.
*/
