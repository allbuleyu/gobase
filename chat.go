package main

import (
	"github.com/allbuleyu/gobase/gopl.io/ch8"
	"net"
	"log"
	"fmt"
)

func main() {
	ch8.Chat()
	return

	listener, _ := net.Listen("tcp", "localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn.RemoteAddr().String())

	}

	return

}
