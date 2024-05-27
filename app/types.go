package main

type Headers map[string]string

type Request struct {
	Method   string
	Protocol string
	Url      string
	Header   Headers
}
