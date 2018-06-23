package ch8

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"time"
)

func Chat() {
	listener, err := net.Listen("udp", "localhost:8080")
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

	tick := time.NewTimer(10 * time.Second)
	go isTimeOutCloseCoon(tick, conn)


	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()

		// 刷新定时器
		tick.Reset(10 * time.Second)
		//tick = time.NewTicker(10 * time.Second)
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func isTimeOutCloseCoon(t *time.Timer, conn net.Conn) {
	for {
		select {
		case <- t.C:
			fmt.Fprintf(conn, "timeout u out %v", t)
			conn.Close()
		}
	}
}

func clientWriter(conn net.Conn, ch <- chan string)  {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

