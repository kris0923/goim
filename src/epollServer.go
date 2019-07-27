package main

import "C"
import (
	//	"bufio"
	"fmt"
	"log"
	//	"net"
//	"syscall"
)

const MAX_EVENTS = 500

func main() {
	// Free BSD create Kqueue
	g_epollFd, err := syscall.EpollCreate(MAX_EVENTS)

}

//type client chan<- string // 对外发送消息的通道
//var (
//	entering = make(chan client)
//	leaving  = make(chan client)
//	message  = make(chan string)
//)

//func boardcaster2() {
//	cliens := make(map[client]bool)
//	epfd, err := syscall.EpollCreate(256)
//	if err != nil {
//		log.Fatalln(err)
//		return false
//	}
//	var et = syscall.EpollEvent
//	et.Events = EPOLLIN | EPOLLET
//	rt.Fd = listen_sock
//	syscall.EpollCtl(epfd, EPOLL_CTL_ADD, listen_sock)
//}

//func handleConn(conn net.Conn) {
//	ch := make(chan string) //  发送客户消息的通道
//	go clientWriter(conn, ch)

//	who := conn.RemoteAddr().String()
//	ch <- "You are " + who
//	message <- who + "has arrived"
//	entering <- ch

//	input := bufio.NewScanner(conn)
//	for input.Scan() {
//		message <- who + ":" + input.Text()
//	}
//	leaving <- ch
//	message <- who + "has left"
//	conn.Close()
//}

//func clientWriter(conn net.Conn, ch <-chan string) {
//	for msg := range ch {
//		fmt.Fprintln(conn, msg)
//	}
//}
