package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	print("here")
	var buf bytes.Buffer
	conn := CreateServer(4421)

	req, err := http.NewRequest("GET", "http//localhost:4221", nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Write(&buf)
	conn.Write(buf.Bytes())

	HandleRequest(conn)

	defer conn.Close()

}
