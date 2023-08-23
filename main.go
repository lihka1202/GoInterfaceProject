package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	readFile(os.Args[1])
}

func readFile(name string) {
	file, err := os.Open(name) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)

}
