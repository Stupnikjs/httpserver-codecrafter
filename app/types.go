package main

import (
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

func NewResponse(statuscode int) *Response {

	response := &Response{
		statuscode: statuscode,
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
	for _, header := range resp.headers {
		sb.WriteString(header)
		sb.WriteString("\r\n")
	}
	sb.WriteString("\r\n")
	sb.WriteString(resp.body)
	return sb.String()
}
