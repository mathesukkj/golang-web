package main

import (
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

		io.WriteString(conn, "teste da coneeeexao\n")
		conn.Close()
	}

}
