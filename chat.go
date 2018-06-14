package main

import (
	"github.com/allbuleyu/gobase/gopl.io/ch8"
	"net"
	"log"
	"fmt"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn.RemoteAddr().String())

	}

	return
	ch8.Chat()
}
