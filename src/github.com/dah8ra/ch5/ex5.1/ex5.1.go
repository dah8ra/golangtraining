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
	//	for _, link := range visit(nil, doc) {
	for _, link := range visits(nil, doc, nil, 0) {
		fmt.Println(link)
	}
}

func visits(links []string, n *html.Node, a []html.Attribute, i int) []string {
	if n == nil {
		return links
	}
	if n.NextSibling != nil {
		links = visits(links, n.NextSibling, n.NextSibling.Attr, 0)
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		if n.Attr[i].Key != "href" {
			links = visits(links, n, n.Attr, i+1)
		} else {
			links = append(links, n.Attr[i].Val)
		}
	}
	links = visits(links, n.FirstChild, n.Attr, 0)

	return links
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
