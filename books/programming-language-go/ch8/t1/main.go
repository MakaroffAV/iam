package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	c := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range c {
				cli <- msg
			}
		case cli := <-entering:
			c[cli] = true
		case cli := <-leaving:
			delete(c, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- ch

	inpit := bufio.NewScanner(conn)
	for inpit.Scan() {
		messages <- who + ": " + inpit.Text()
	}
	leaving <- ch
	messages <- who + " отключился"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	l, lErr := net.Listen("tcp", "localhost:8000")
	if lErr != nil {
		panic(lErr)
	}

	go broadcaster()

	for {
		c, cErr := l.Accept()
		if cErr != nil {
			log.Print(cErr)
			continue
		}
		go handleConn(c)
	}
}
