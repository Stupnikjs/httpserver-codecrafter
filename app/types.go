package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Headers map[string]string

type Request struct {
	method  string
	version string
	path    string
	headers Headers
}

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
		sb.WriteString(" NOT FOUND\r\n")
	}
	for key, header := range resp.headers {
		sb.WriteString(fmt.Sprintf("%s: %s", key, header))
		sb.WriteString("\r\n")
	}
	sb.WriteString("Content-Type: application/x-www-form-urlencoded")
	sb.WriteString("\r\n")
	sb.WriteString(resp.body)
	return sb.String()
}
