package main

import (
	"log"
	"net"
)
import "goim/server"

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8900")
	if err != nil {
		log.Fatal(err)
	}
	go server.Boardcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go server.HandleConn(conn)
	}
}
