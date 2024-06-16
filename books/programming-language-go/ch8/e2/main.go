package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, wErr := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if wErr != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	l, lErr := net.Listen("tcp", "localhost:8000")
	if lErr != nil {
		panic(lErr)
	}

	for {
		c, cErr := l.Accept()
		if cErr != nil {
			log.Print(cErr)
			continue
		}

		// без go будет последовательная
		// версия              программы
		go handleConn(c)
	}

}
