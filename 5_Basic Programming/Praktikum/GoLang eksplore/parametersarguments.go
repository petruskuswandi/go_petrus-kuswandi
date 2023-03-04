package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func familysName(fnames string, age int) {
	fmt.Println("Hello", age, "year old", fnames, "Refsnes")
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")

	fmt.Printf("\n")

	familysName("Liam", 3)
	familysName("Jenny", 14)
	familysName("Anja", 30)
}
