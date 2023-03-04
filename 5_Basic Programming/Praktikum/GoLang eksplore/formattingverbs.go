package main

import (
	"fmt"
)

func main() {
	var i = 15.5
	var txt = "Hello World"
	//General Formatting verbs

	fmt.Printf("%v\n", i)   // Prints the value in the default format
	fmt.Printf("%#v\n", i)  // Prints the value in Go-syntax format
	fmt.Printf("%T\n", i)   // Prints the type of the value
	fmt.Printf("%v%%\n", i) // Prints the % sign

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	fmt.Printf("\n")

	//Integer Formatting Verbs
	var j = 15

	fmt.Printf("%b\n", j)   // Base 2
	fmt.Printf("%d\n", j)   // Base 10
	fmt.Printf("%+d\n", j)  // Base 10 adn always show sign
	fmt.Printf("%o\n", j)   // Base 8
	fmt.Printf("%O\n", j)   // Base 8, with leading 0o
	fmt.Printf("%x\n", j)   // Base 16, lowercase
	fmt.Printf("%X\n", j)   // Base 16, uppercase
	fmt.Printf("%#x\n", j)  // Base 16, with leading 0x
	fmt.Printf("%4d\n", j)  // Pad with the spaces (width 4, right justified)
	fmt.Printf("%-4d\n", j) // Pad with the spaces (width 4, left justified)
	fmt.Printf("%04d\n", j) // Pad with zeroes (width 4)

	fmt.Printf("\n")

	var text = "hello"

	fmt.Printf("%s\n", text)   // Prints the value as plain string
	fmt.Printf("%q\n", text)   // Prints the values as a double-quoted string
	fmt.Printf("%8s\n", text)  // Prints the value as a plain string (width 8, right justified)
	fmt.Printf("%-8s\n", text) // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", text)   // Prints the value as hex dump of byte values
	fmt.Printf("% x\n", text)  // Prints the value as hex dump with spaces

	fmt.Printf("\n")

	var k = true
	var l = false

	fmt.Printf("%t\n", k) // Value of the boolean operator in true or false format (same as using %v)
	fmt.Printf("%t\n", l)

	fmt.Printf("\n")

	var m = 3.141
	fmt.Printf("%e\n", m)    // Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", m)    // Decimal point, no exponent
	fmt.Printf("%.2f\n", m)  // Default width, precision 2
	fmt.Printf("%6.2f\n", m) // Width, precision 2
	fmt.Printf("%g\n", m)    // Exponent as needed, only necessary digits
}
