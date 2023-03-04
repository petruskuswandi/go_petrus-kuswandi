package main

import (
	"fmt"
)

func main() {
	var a = 15 + 25
	fmt.Println(a)

	fmt.Printf("\n")

	var (
		sum1 = 100 + 50
		sum2 = sum1 + 25
		sum3 = sum2 + sum2
	)
	fmt.Println(sum3)

	fmt.Printf("\n")

	// Arithmetic
	x := 10
	y := 2
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	x++
	fmt.Println(x)
	y--
	fmt.Println(y)

	fmt.Printf("\n")

	//Assignment Operators
	var z = 5
	fmt.Println(z)

	z += 3 // z = z + 3
	fmt.Println(z)

	z -= 3
	fmt.Println(z)

	z *= 3
	fmt.Println(z)

	z /= 3
	fmt.Println(z)

	z %= 3
	fmt.Println(z)

	var g = 5
	fmt.Printf("g is %b \n", g)
	fmt.Printf("3 is %03b \n", 3)

	g &= 3
	fmt.Printf("g now is %03b \n", g)

	var h = 5
	fmt.Printf("h is %b \n", h)
	fmt.Printf("3 is %03b \n", 3)

	h |= 3
	fmt.Printf("h now is %03b \n", h)

	var i = 5
	fmt.Printf("i is %b \n", i)
	fmt.Printf("3 is %03b \n", 3)

	i ^= 3
	fmt.Printf("i now is %03b \n", i)

	var j = 5
	fmt.Printf("f is %b \n", j)

	j >>= 3
	fmt.Printf("j now is %03b \n", j)

	var k = 5
	fmt.Printf("k is %b \n", k)

	k <<= 3
	fmt.Printf("k is %b \n", k)

	fmt.Printf("\n")

	//Comparrisons
	var e = 5
	var f = 3
	fmt.Println(e > f)
	fmt.Println(e < f)
	fmt.Println(e == f)
	fmt.Println(e != f)
	fmt.Println(e >= f)
	fmt.Println(e <= f)

	fmt.Printf("\n")

	//Logical
	var q = 5

	fmt.Println(q < 5 && q < 10)
	fmt.Println(q < 5 || q < 4)
	fmt.Println(!(q < 5 && q < 10))

	fmt.Printf("\n")

	// Bitwise
	var b = 9
	var c = 8

	fmt.Printf("b = %b\n", b)
	fmt.Printf("c = %b\n", c)

	fmt.Printf("b & c is %b\n", b&c)
	fmt.Printf("b & c is %b\n", b|c)
	fmt.Printf("b & c is %b\n", b^c)

	var d = 9
	fmt.Printf("d = %b\n", d)
	fmt.Printf("d << 2 is %b\n", d<<2)
	fmt.Printf("d >> 2 is %b\n", d>>2)
}
