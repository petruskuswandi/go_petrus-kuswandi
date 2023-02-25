package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
Konstanta
Konstanta dideklarasikan seperti variabel, tapi dengan kata const.

Konstanta bisa bernilai karakter, string, boolean, atau numerik.

Konstanta tidak bisa dideklarasikan dengan sintaks :=.
*/
