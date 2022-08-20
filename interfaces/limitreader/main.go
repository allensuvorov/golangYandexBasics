package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func LimitReader(r io.Reader, n int) io.Reader {
	// ...
	b := make([]byte, n)
	r.Read(b)
	return strings.NewReader(string(b) + "\n")

}

func main() {
	// new struct with string and Reader method inside
	r := strings.NewReader("some io.Reader stream to be read\n")
	// call LimitReader function for the Reader struct and limit
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
