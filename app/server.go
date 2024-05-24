package main

import (
	"fmt"
	"net"
	"os"
)

func HandleRequest(conn net.Conn) {
	// Read the request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer))

}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	conn := CreateServer(4222)
	defer conn.Close()

	HandleRequest(conn)
}

func CreateServer(port int) net.Conn {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		srt := fmt.Sprintf("Failed to bind to port %d", port)
		print(srt)
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Failed to accept")
		os.Exit(1)
	}
	return conn

}
