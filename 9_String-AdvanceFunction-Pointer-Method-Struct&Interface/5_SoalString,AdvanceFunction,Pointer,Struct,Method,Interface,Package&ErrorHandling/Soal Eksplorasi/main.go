package main

import "fmt"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type substitutionCipher struct {
	keyMap map[rune]rune
}

func (c substitutionCipher) Encode(text string) string {
	encodedText := ""
	for _, char := range text {
		if val, ok := c.keyMap[char]; ok {
			encodedText += string(val)
		} else {
			encodedText += string(char)
		}
	}
	return encodedText
}

func (c substitutionCipher) Decode(text string) string {
	reverseMap := make(map[rune]rune)
	for key, val := range c.keyMap {
		reverseMap[val] = key
	}
	decodedText := ""
	for _, char := range text {
		if val, ok := reverseMap[char]; ok {
			decodedText += string(val)
		} else {
			decodedText += string(char)
		}
	}
	return decodedText
}

func main() {
	var menu int
	var a string
	var c Cipher = substitutionCipher{
		keyMap: map[rune]rune{
			'a': 'z',
			'b': 'y',
			'c': 'x',
			'd': 'w',
			'e': 'v',
			'f': 'u',
			'g': 't',
			'h': 's',
			'i': 'r',
			'j': 'q',
			'k': 'p',
			'l': 'o',
			'm': 'n',
			'n': 'm',
			'o': 'l',
			'p': 'k',
			'q': 'j',
			'r': 'i',
			's': 'h',
			't': 'g',
			'u': 'f',
			'v': 'e',
			'w': 'd',
			'x': 'c',
			'y': 'b',
			'z': 'a',
		},
	}

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Text: ")
		fmt.Scan(&a)
		fmt.Print("\nEncoded text of " + a + " is: " + c.Encode(a))
	} else if menu == 2 {
		fmt.Print("\nInput Encoded Text: ")
		fmt.Scan(&a)
		fmt.Print("\nDecoded text of " + a + " is: " + c.Decode(a))
	}
}
