package main

import (
	"Clock-server/handlers"
	"flag"
	"fmt"
	"log"
	"net"
)

var f *uint = flag.Uint("f", 100, "updating time frequency in Millisecond as a unit")
var c *bool = flag.Bool("c", false, "use a concurent connections")

func init() {
	flag.Parse()
	handlers.Frequency = *f
}

func main() {
	server, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	fmt.Println("server started,waiting for connections")
	if *c {
		handlers.ConcurentHandling(server)
	} else {
		handlers.SequentialHandling(server)
	}
}
