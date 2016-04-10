package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	fmt.Printf("============\n%s\n============\n", p)

	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*wc += WordCounter(count)

	return len(p), nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	fmt.Printf("============\n%s\n============\n", p)
	count := 0
	for _, v := range p {
		if v == '\n' {
			count++
		}
	}
	*lc += LineCounter(count + 1)

	return len(p), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello world test"))
	fmt.Println("Word: ", wc)

	fmt.Println("---------------------")

	var lc LineCounter
	lc.Write([]byte("hello\nworld\ntest"))
	fmt.Println("Line: ", lc)
}
