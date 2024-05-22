package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

func HandleRequest(conn net.Conn) {
	// Read the request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	reader := bufio.NewReader(bytes.NewBuffer(buffer))

	// reads until '\n'
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// split fields by " "
	requestFields := strings.Fields(requestLine)
	if requestFields[1] == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	conn := CreateServer()

	defer conn.Close()
	HandleRequest(conn)
}

func CreateServer() net.Conn {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
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
