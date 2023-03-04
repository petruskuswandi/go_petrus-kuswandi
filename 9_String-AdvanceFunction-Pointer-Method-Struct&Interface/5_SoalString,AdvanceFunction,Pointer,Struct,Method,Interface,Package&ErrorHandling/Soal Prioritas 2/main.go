package main

import "fmt"

func caesar(offset int, input string) string {
	var output string
	for _, c := range input {
		if c == ' ' {
			output += " "
		} else {
			shifted := rune(int(c-'a'+rune(offset))%26 + int('a'))
			output += string(shifted)
		}
	}
	return output
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
