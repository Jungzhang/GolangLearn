package main

import (
	"net"
	"log"
	"fmt"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 1)
	time.Sleep(time.Second * 2)
	for {
		fmt.Println("aaaa")
		n, err := conn.Write(b)
		if err != nil {
			log.Println(err)
			break
		} else {
			fmt.Println("n =", n)
		}
	}
}
