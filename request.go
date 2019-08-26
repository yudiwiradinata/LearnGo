package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	//bs := make([]byte, 9999999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("print size", len(b))
	return len(b), nil
}
