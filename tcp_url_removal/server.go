package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			xs := strings.Split(ln, " ")
			path := xs[1]
			fmt.Printf("localhost:5050%s\r\n", path)
		}

		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	fmt.Fprint(conn, "oi")
}
