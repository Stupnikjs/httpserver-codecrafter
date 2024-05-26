package main

type Request struct {
	Method   string
	Protocol string
	Url      string
	Headers  map[string]string
}
