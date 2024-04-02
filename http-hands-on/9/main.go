package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		go serve(conn)

	}
}

func serve(conn net.Conn) {
	sc := bufio.NewScanner(conn)

	for sc.Scan() {
		ln := sc.Text()
		if ln == "" {
			break
		}

		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(ln))
		fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	}
	conn.Close()
}
