package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8900")
	if err != nil {
		conn.Close()
		fmt.Println("conn error")
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("read error")
		}
		fmt.Println("status is:" + status)
	}

}
