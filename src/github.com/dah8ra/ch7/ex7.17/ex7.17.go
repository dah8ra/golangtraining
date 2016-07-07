package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// Sample command
// go run ex7.17.go css.xml id
func main() {
	dec := xml.NewDecoder(getFileContents())
	//	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	temp := ""
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		flag := false
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push start element
			if len(tok.Attr) > 0 {
				for i := 0; i < len(tok.Attr); i++ {
					temp = tok.Attr[i].Name.Local
					stack = append(stack, tok.Attr[i].Name.Local) // push attribute
					if containsAll(stack, os.Args[2:]) {
						fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok.Attr[i].Value)
					}
					stack = stack[:len(stack)-1] // pop attribute
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[2:]) {
				for _, ss := range stack {
					if strings.Contains(ss, temp) {
						flag = true
						break
					}
				}
				if !flag { // Remove useless print output
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
			}
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
