package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	jsonVar := bytes.NewBuffer([]byte(`{"name": "akon"}`))
	c := http.Client{Timeout: time.Second}
	resp, err := c.Post("http://google.com.br", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
