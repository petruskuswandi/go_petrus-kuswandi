package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
Kondisi ("if")
Perintah if mirip seperti pada pengulangan for; ekspresinya tidak harus ditutupi dengan tanda-kurung ( ) namun tanda { } diharuskan.
*/
