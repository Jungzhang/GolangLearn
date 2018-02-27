package main

import (
	"fmt"
	"time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		i := 0
		for i <= 5 {
			time.Sleep(time.Second)
			naturals <- i
			i++
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		x, ok := <-squares
		if ok {
			fmt.Println(x)
		} else {
			break
		}
	}
}
