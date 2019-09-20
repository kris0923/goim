package server

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string // 对外发送消息的通道
var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func Boardcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func HandleConn(conn net.Conn) {
	ch := make(chan string) //  发送客户消息的通道
	go ClientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ":" + input.Text()
	}
	leaving <- ch
	message <- who + "has left"
	conn.Close()
}

func ClientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
