package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":4080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)

	var i int
	var method, uri string

	for sc.Scan() {
		ln := sc.Text()

		if i == 0 {
			xs := strings.Fields(ln)
			method = xs[0]
			uri = xs[1]
		}

		if ln == "" {
			break
		}

		i++
	}
	body := "body payload\n"
	body += method
	body += "\n"
	body += uri
	body += "\n"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Println(body)
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
