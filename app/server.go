package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
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
	ReadResp(conn)
	//conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

}

func ReadResp(conn net.Conn) {
	str := ""
	for {
		buf := make([]byte, 1)
		_, err := conn.Read(buf)

		if err == io.EOF {
			break
		}
		str += string(buf)

	}
	print(str)

}
