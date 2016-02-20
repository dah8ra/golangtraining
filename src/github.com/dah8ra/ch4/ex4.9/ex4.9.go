package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fmt.Printf(">> read file: %s\n", os.Args[1])
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		scanner := bufio.NewScanner(strings.NewReader(string(line)))
		scanner.Split(bufio.ScanWords)
		count := 0
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		fmt.Printf("%d\n", count)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
