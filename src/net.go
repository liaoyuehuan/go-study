package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {

	conn, err := net.Dial("tcp", "192.168.50.74:9501")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("status:", status)
}
