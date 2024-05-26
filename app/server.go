package main

import (
	"fmt"
	"io"

	// Uncomment this block to pass the first stage
	"bufio"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	defer conn.Close()
	handleConn(conn)
}
func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error reading request: ", err.Error())
		return
	}
	url := strings.Split(request, " ")[1]
	if url == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
