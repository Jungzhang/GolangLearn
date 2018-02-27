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
		for {
			time.Sleep(time.Second)
			i++
			naturals <- i
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}

	}()

	for {
		for {
			fmt.Println(<-squares)
		}
	}
}
