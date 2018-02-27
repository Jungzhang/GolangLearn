package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			ch1 <- 1
		}

	}()

	for {
		select {
		case x := <-ch1:
			fmt.Println(x)
		default:
			fmt.Println("default")
		}
	}

}
