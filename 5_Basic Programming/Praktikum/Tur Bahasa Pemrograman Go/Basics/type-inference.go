package main

import "fmt"

func main() {
	v := 42 // ubahlah nilai v!
	fmt.Printf("v bertipe %T\n", v)
}

/*
Inferensi tipe
Saat mendeklarasikan sebuah variabel tanpa menentukan tipe eksplisitnya (baik menggunakan sintaks := atau var), tipe variabel disamakan dengan nilai di sebelah kanannya.

Bila sisi kanan dari deklarasi ditulis tipenya, maka variabel baru memiliki tipe yang sama:

var i int
j := i // j adalah sebuah int
Tapi bila sisi kanan berupa konstanta numerik, variabel baru bisa jadi int, float64, atau complex128 bergantung kepada presisi dari konstanta:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
Coba ubah nilai awal dari v pada contoh kode dan perhatikan bagaimana tipenya berubah.


*/
