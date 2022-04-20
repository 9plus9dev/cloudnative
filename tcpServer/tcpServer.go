package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ... ")
	listener, err := net.Listen("tcp", ":5080")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting", err.Error())
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		fmt.Printf("the client sent data have length %v\n", len)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
