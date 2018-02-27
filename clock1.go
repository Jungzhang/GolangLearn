package main

import (
	"net"
	"log"
	"io"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // 奇葩的go语言时间格式化=_=
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
