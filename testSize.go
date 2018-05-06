package main

import (
	"fmt"
	"unsafe"
)

const (
	A = "aaaaa"
	B = "bbbbbbbb"
	C = "ccc"
	D = "D"
)

const (
	AA = iota
	BB
	CC
	DD
)

func main() {

	fmt.Printf("A = %d, B = %d, C = %d, D = %d\nAA = %d, BB = %d, CC = %d, DD = %d\n",
		unsafe.Sizeof(A), unsafe.Sizeof(B), unsafe.Sizeof(C), unsafe.Sizeof(D),
			unsafe.Sizeof(AA), unsafe.Sizeof(BB), unsafe.Sizeof(CC), unsafe.Sizeof(DD))

}
