package main

import "fmt"

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil", "akume"}, []string{"eddie", "steve", "geese"}))
	//["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}

func ArrayMerge(arrayA []string, arrayB []string) []string {
	mergedSlice := append(arrayA[:], arrayB...) // Copy elemen dari slice1 ke mergedSlice dan tambahkan elemen dari slice2
	addedNames := make(map[string]bool)         // Buat map untuk menyimpan nama yang sudah ditambahkan
	resultSlice := []string{}                   // Buat slice untuk menampung elemen yang unik

	// Loop setiap elemen di mergedSlice
	for _, name := range mergedSlice {
		if _, ok := addedNames[name]; !ok {
			addedNames[name] = true                 // Tandai nama yang sudah ditambahkan
			resultSlice = append(resultSlice, name) // Tambahkan nama ke resultSlice
		}
	}
	return resultSlice
}
