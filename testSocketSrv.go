package main

import (
	"net"
	"log"
	"fmt"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("a new client connected")
			conn.Close()
			fmt.Println("a client closed")
		}(conn)
	}
}
