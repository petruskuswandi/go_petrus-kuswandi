package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

/*
Pointer ke struct
Bagian dari struct dapat diakses lewat pointer ke struct.

Untuk mengakses bagian X dari sebuah struct bila kita memiliki pointer ke struct p, kita dapat menulisnya dengan (*p).X. Namun, notasi tersebut tidak praktis, sehingga Go membolehkan kita mengaksesnya langsung dengan menulis hanya p.X.
*/
