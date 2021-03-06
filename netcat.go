package main

import (
	"net"
	"log"
	"os"
	"io"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
