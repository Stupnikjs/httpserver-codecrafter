package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	var buf, respBuf bytes.Buffer
	conn := CreateServer()

	req, err := http.NewRequest("GET", "http//localhost:4221", nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Write(&buf)
	conn.Write(buf.Bytes())

	conn.Read(respBuf.Bytes())
	fmt.Println(respBuf.String())

}
