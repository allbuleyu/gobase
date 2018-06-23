package ch8

import (
	"time"
	"net"
	"fmt"
)

type c struct {
	addr string
	port int64
	name string
	lastTime time.Time
}

type client1 chan <- c

var (
	enter = make(chan client1)
	leave = make(chan client1)
	msg = make(chan string)
)

func Chat1() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err !=nil {
		fmt.Errorf("connect err: %v", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Errorf("accept err: %v", err)
	}

	go connHandle(conn)

	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Fprintf(os.Stdout, "accept err: %v", err)
	//		continue
	//	}
	//
	//
	//}

	<- msg
}


func connHandle(conn net.Conn) {
	//cli := c{}
	//cli.addr = conn.RemoteAddr().String()
	//cli.name = cli.addr
	//cli.lastTime = time.Now()
	//
	//cc := make(chan c)
	//cc <- cli
	tick := time.Tick(5*time.Second)

	i := 0
	for e := range tick {
		fmt.Fprintf(conn, "%v \n", e)
		i++
		if i == 5 {
			break
		}
	}

	msg <- "done sss"
}