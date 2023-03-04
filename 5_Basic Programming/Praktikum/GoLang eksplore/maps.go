package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	var c = make(map[string]string)
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)

	var e = make(map[string]string)
	var f map[string]string

	fmt.Println(e == nil)
	fmt.Println(f == nil)

}
