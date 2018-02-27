package main

import (
	"net"
	"log"
	"bufio"
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
		go handerConn(conn)
	}
}

func handerConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
	}
	conn.Close()
}
