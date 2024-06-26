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

		sc := bufio.NewScanner(conn)

		for sc.Scan() {
			fmt.Println(sc.Text())
		}

		fmt.Println("code got here")
		io.WriteString(conn, "I see you connected rsrsrs")
	}
}
