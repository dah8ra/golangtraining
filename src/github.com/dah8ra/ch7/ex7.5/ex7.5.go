package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	r := bytes.NewReader([]byte("Hello\n"))
	LimitReader(r, 5)
	LimitReader(r, 5)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	b := make([]byte, n)
	num, err := io.ReadFull(r, b)
	if err != nil {
		fmt.Printf("Reached EOF. The size is %d bytes.\n", num)
	} else {
		fmt.Printf("Remain some data in the buffer. The reading size is %d bytes.\n", num)
	}

	return r
}
