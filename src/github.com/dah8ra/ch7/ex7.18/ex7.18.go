package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Sample command
// go run ex7.18.go test.xml
func main() {
	dec := xml.NewDecoder(getFileContents())
	//	dec := xml.NewDecoder(os.Stdin)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Printf("%s", tok.Name)
			for _, attr := range tok.Attr {
				fmt.Printf("%s", attr)
			}
		case xml.EndElement:
			fmt.Printf("%s", tok.Name)
		case xml.CharData:
			fmt.Printf("%s", tok)
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	//	fmt.Printf("x: %v\ty: %v\n", x, y)
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
			//			fmt.Printf("y: %v\n", y)
		}
		x = x[1:]
		//		fmt.Printf("x: %v\n", x)
	}
	return false
}

func getFileContents() *bufio.Reader {
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
	}

	return bufio.NewReaderSize(fp, 4096)
}
