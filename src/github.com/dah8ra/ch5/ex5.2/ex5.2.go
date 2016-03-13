package main

import (
	"bufio"
	"fmt"
	"os"

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
	m := make(map[string]int)
	for key, value := range visit(m, doc) {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}

func visit(m map[string]int, n *html.Node) map[string]int {
	fmt.Printf("%s\n", n.Data)
	if n.Type == html.ElementNode {
		if n.Data == "p" {
			m["p"]++
		} else if n.Data == "div" {
			m["div"]++
		} else if n.Data == "span" {
			m["span"]++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = visit(m, c)
	}
	return m
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
