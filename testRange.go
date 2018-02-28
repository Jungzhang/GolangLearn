package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 3; i++ {
		ch <- "HelloWorld"
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {

	ch := make(chan string)

	go sendMessage(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
