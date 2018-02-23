package main

import "fmt"

func main() {
	i, j := 1, 2
	fmt.Println("i = ", i, "   j = ", j)
	i, j = j, i
	fmt.Println("i = ", i, "  j = ", j)
}
