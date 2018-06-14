package ch8

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func Chat() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}


type client chan <- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 向所有在线用户发消息
			for cli := range clients {
				cli <- msg
			}
		case cli := <- entering:
			// 新进用户
			clients[cli] = true
		case cli := <- leaving:
			// 用户离开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {

	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string)  {
	for msg := range messages {

		fmt.Fprintln(conn, msg)
	}
}

