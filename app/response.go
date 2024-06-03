package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	statuscode int
	body       string
	headers    Headers
}

func NewResponse(statuscode int, headers Headers) *Response {
	response := &Response{
		statuscode: statuscode,
		headers:    headers,
	}
	return response
}

func (resp *Response) addTextBody(str string) *Response {
	resp.body = str
	return resp
}

func (resp *Response) toString() string {
	var sb strings.Builder
	sb.WriteString("HTTP/1.1 ")
	status := strconv.Itoa(resp.statuscode)

	sb.WriteString(status)
	if status == "200" {
		sb.WriteString(" OK\r\n")
	}
	if status == "404" {
		sb.WriteString(" Not Found\r\n")
	}
	for key, header := range resp.headers {
		sb.WriteString(fmt.Sprintf("%s: %s", key, header))
		sb.WriteString("\r\n")
	}
	sb.WriteString("Content-Type: text/plain")
	sb.WriteString("\r\n")
	sb.WriteString(fmt.Sprintf("Content-Length: %d", len(resp.body)))
	sb.WriteString("\r\n")
	sb.WriteString("\r\n")
	sb.WriteString(resp.body)
	return sb.String()
}
