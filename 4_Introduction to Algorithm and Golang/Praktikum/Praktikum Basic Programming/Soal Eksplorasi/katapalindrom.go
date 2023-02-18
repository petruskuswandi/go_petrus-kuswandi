package main

import (
	"bufio"
	"fmt"
	"os"
)

func palindromo(word string) bool {

	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-1-i] {
			return false
		}

	}

	return true
}

func main() {
	fmt.Print("Masukkan kata atau kalimat yang ingin di cek: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
	result := palindromo(line)
	fmt.Println(result)
	if result == true {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}
}
