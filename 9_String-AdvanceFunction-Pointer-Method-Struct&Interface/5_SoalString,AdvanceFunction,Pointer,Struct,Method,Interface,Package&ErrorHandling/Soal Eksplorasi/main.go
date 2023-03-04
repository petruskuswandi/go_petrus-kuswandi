package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Cipher interface {
	Encode() string
	Decode() string
}

type SubstitutionCipher struct {
	keyMap map[rune]rune
}

func NewSubstitutionCipher(key string) *SubstitutionCipher {
	keyMap := make(map[rune]rune)
	for i, c := range key {
		keyMap[c] = rune('a' + i)
	}
	return &SubstitutionCipher{keyMap}
}

func (s *SubstitutionCipher) Encode(text string) string {
	var result []rune
	for _, c := range text {
		if val, ok := s.keyMap[c]; ok {
			result = append(result, val)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

func (s *SubstitutionCipher) Decode(text string) string {
	var result []rune
	for _, c := range text {
		if val, ok := s.getKeyByValue(c); ok {
			result = append(result, val)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

func (s *SubstitutionCipher) getKeyByValue(value rune) (rune, bool) {
	for k, v := range s.keyMap {
		if v == value {
			return k, true
		}
	}
	return 0, false
}

func (s *student) Encode() string {
	cipher := NewSubstitutionCipher("rizky")
	s.nameEncode = cipher.Encode(s.name)
	return s.nameEncode
}

func (s *student) Decode() string {
	cipher := NewSubstitutionCipher("rizky")
	s.name = cipher.Decode(s.nameEncode)
	return s.name
}

func main() {
	var menu int
	var a student = student{}
	var c Cipher = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Encoded Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
