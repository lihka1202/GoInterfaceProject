package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type newWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}

func (newWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("This was a custom writer function that I made!n I wonder what are some of the other functions that could work this way!")
	return len(bs), nil
}
