package handlers

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var Frequency uint = 100
var activeConn activeConncetion

func init() {
	activeConn = activeConncetion{}
}

func SequentialHandling(server net.Listener) {
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("new connection,", activeConn.Add(1), "active connection")
		clientHandler(conn)
	}
}

func ConcurentHandling(server net.Listener) {
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("new connection,", activeConn.Add(1), "active connection")
		go clientHandler(conn)
	}
}

func clientHandler(conn net.Conn) {
	defer conn.Close()
	for {
		t := "\r" + time.Now().String()
		_, err := io.WriteString(conn, t)
		if err != nil {
			fmt.Println("connection closed,", activeConn.Add(-1), "active connection")
			return
		}
		time.Sleep(time.Duration(Frequency) * time.Millisecond)
	}
}
