package main

import "fmt"

func main() {
	const (
		a = iota
		b// = 3 << iota
		c
		d// = 100
		e
		f
		g
	)

	fmt.Println(a, b, c, d, e, f, g)
}
