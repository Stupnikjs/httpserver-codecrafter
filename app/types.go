package main

import "fmt"

type Headers map[string]string

type Request struct {
	method  string
	version string
	path    string
	headers Headers
}

type Response struct {
	statuscode int
	body       interface{}
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
	responseString := fmt.Sprintf("response %s", "hello")
	return responseString
}
