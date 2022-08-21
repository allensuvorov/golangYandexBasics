package main

import (
	"io"
	"log"
	"os"
	"strings"
)

//  Lireader struct
type LimitedReader struct {
	reader io.Reader
	left   int
}

func LimitReader(r io.Reader, n int) io.Reader {

	return &LimitedReader{reader: r, left: n}
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}
	n, err := r.reader.Read(p)
	r.left -= n
	return n, err
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
