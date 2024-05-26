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
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	ParseRequest(request, conn)

}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	if err != nil {
		fmt.Println("Failed to accept")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to accept")
			os.Exit(1)
		}

		go HandleRequest(conn)

	}

}

func ParseRequest(request string, conn net.Conn) {
	// split fields by " "
	requestFields := strings.Fields(request)
	url := requestFields[1]

	paths := []string{
		"/echo",
		"/",
	}

	for _, route := range paths {
		if _, ok := strings.CutPrefix(url, route); ok {
			respBody := strings.TrimPrefix(url, "/echo/")
			resp := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(respBody), respBody)
			conn.Write([]byte(resp))
			return
		}
	}

	respBody := ""
	resp := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(respBody), respBody)
	conn.Write([]byte(resp))

}
