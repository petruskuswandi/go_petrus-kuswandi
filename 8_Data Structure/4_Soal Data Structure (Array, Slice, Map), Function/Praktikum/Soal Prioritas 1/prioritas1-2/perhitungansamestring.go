package main

import "fmt"

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:2]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))               // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))
}

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int) // Buat map untuk menyimpan hitungan kemunculan setiap string

	// Loop setiap elemen di slice
	for _, element := range slice {
		countMap[element]++ // Tambahkan hitungan kemunculan untuk string tersebut di map
	}

	return countMap
}
