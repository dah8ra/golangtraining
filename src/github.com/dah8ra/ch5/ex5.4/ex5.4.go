package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	reader := getFileContents()
	//	doc, err := html.Parse(os.Stdin)
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if strings.Contains(n.Data, "script") {
		for _, v := range n.Attr {
			fmt.Printf("DETECTED script -> %s\n", v)
		}
	}
	if strings.Contains(n.Data, "img") {
		for _, v := range n.Attr {
			fmt.Printf("DETECTED img -> %s\n", v)
		}
	}
	if strings.Contains(n.Data, "div") {
		for _, a := range n.Attr {
			if a.Key == "style" {
				fmt.Printf("DETECTED style -> %s\n", a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
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
		//		defer fp.Close()
	}

	return bufio.NewReaderSize(fp, 4096)
}
