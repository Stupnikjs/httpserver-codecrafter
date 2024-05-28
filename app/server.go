package main

import (
	"fmt"

	// Uncomment this block to pass the first stage
	"bufio"
	"net"
	"os"
	"strings"
)

func HandleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection received")
	s := bufio.NewScanner(conn)
	s.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for s.Scan() {
		if text := s.Text(); text != "" {
			lines = append(lines, text)
		} else {
			break
		}
	}
	fmt.Println("Request:", strings.Join(lines, ", "))
	req := parseRequest(lines)
	pathParts := strings.Split(req.path, "/")[1:]
	var response *Response
	switch pathParts[0] {
	case "":
		response = NewResponse(200)
	case "echo":
		response = NewResponse(200).addTextBody(strings.Join(pathParts[1:], "/"))
	case "user-agent":
		response = NewResponse(200).addTextBody(req.headers["User-Agent"])
	default:
		response = NewResponse(404)
	}

	resStr := response.toString()
	fmt.Println("Responding:", resStr)
	conn.Write([]byte(resStr))
}

func parseRequest(lines []string) *Request {
	startLine := strings.Split(lines[0], " ")
	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if line != "" {
			split := strings.Split(line, ": ")
			headers[split[0]] = split[1]
		}
	}
	return &Request{
		method:  startLine[0],
		path:    startLine[1],
		version: startLine[2],
		headers: headers,
	}
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
