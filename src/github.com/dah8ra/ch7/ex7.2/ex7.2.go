package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := []byte("hello world")
	writer := bufio.NewWriter(w)
	n, err := writer.Write(b)
	if err != nil {
		panic(err)
	}
	var i64 int64
	i64 = int64(n)
	fmt.Printf("Words  : %d\n", n)
	return writer, &i64
}

func main() {
	_, c := CountingWriter(os.Stdout)
	fmt.Printf("Pointer: %d\n", c)
}
