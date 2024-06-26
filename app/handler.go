package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Headers map[string]string

type Request struct {
	method  string
	version string
	path    string
	headers Headers
}

func HandleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection received")
	s := bufio.NewReader(conn)

	lines := make([]string, 0)
	for {
		line, _, _ := s.ReadLine()
		if string(line) != "" {
			lines = append(lines, string(line))
		} else {
			break
		}
	}

	req := parseRequest(lines)
	pathParts := strings.Split(req.path, "/")[1:]
	var response *Response
	switch pathParts[0] {
	case "":
		response = NewResponse(200, req.headers)
	case "echo":
		response = NewResponse(200, req.headers).addTextBody(strings.Join(pathParts[1:], "/"))
	case "user-agent":
		response = NewResponse(200, req.headers).addTextBody(req.headers["User-Agent"])
	default:
		response = NewResponse(404, req.headers)
	}

	resStr := response.toString()

	conn.Write([]byte(resStr))
}

func parseRequest(lines []string) *Request {

	startLine := strings.Split(lines[0], " ")

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		fmt.Println(line)
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
