package main

import (
	"bytes"
	"fmt"

	// Uncomment this block to pass the first stage
	"bufio"
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

	headerStr, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	headers := ParseHeaderStr(headerStr)

	ParseRequest(request, headers, conn)

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
		fmt.Println("Error accepting connection: ", err.Error())
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

func ParseRequest(request string, headers Headers, conn net.Conn) Request {
	// split fields by " "
	req := Request{}

	fmt.Println(headers)
	requestFields := strings.Fields(request)
	req.Method = requestFields[0]
	req.Url = requestFields[1]
	req.Protocol = requestFields[2]
	fmt.Println(requestFields)

	if req.Url == "/" {
		head := Headers{
			"Content-type":   "text/plain",
			"Content-Length": "0",
		}
		resp := fmt.Sprintf("HTTP/1.1 200 OK\r\n%s\r\n%s", head.ToString(), "")
		conn.Write([]byte(resp))
		return req
	}

	if req.Url == "/user-agent" {
		head := Headers{
			"Content-type":   "text/plain",
			"Content-Length": "0",
			"User-Agent":     "",
		}
		resp := fmt.Sprintf("HTTP/1.1 404 Not Found\r\n%s\r\n%s", head.ToString(), "")
		conn.Write([]byte(resp))
		return req
	}
	respBody := ""
	head := Headers{
		"Content-type":   "text/plain",
		"Content-Length": "0",
	}

	resp := fmt.Sprintf("HTTP/1.1 404 Not Found\r\n%s\r\n%s", head.ToString(), respBody)
	conn.Write([]byte(resp))

}

func (h Headers) ToString() string {
	str := ""
	for k, v := range h {
		str += fmt.Sprintf("%s: %s", k, v)
	}
	return str + "\r\n"
}

func ParseHeaderStr(str string) Headers {
	fields := strings.Split(str, "/r/n")

	fmt.Println(fields)
	return nil
}
