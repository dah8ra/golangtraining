package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var id string = "script"

func main() {
	reader := getFileContents()
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	ElementByID(doc, id)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	if ok := forEachNode(doc, startElement, endElement); ok == false {
		fmt.Println("--- Terminate processing ---\n")
	}
	return doc
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool) bool {
	if pre != nil {
		if ok := pre(n, id); ok == false {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if ok := forEachNode(c, pre, post); ok == false {
			return false
		}
	}
	displayAttribute(n)

	if post != nil {
		if ok := post(n, id); ok == false {
			return false
		}
	}
	return true
}

func displayAttribute(n *html.Node) {
	for _, a := range n.Attr {
		fmt.Printf("%*s%s\n", depth*2, "", a.Val)
	}
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Data == id {
		fmt.Printf("@@@ DETECTED -> %s\n", n.Data)
		return false
	}
	if n.Type == html.ElementNode {
		if len(n.Attr) != 0 {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
	return true
}

func endElement(n *html.Node, id string) bool {
	if n.Data == id {
		fmt.Printf("@@@ DETECTED -> %s\n", n.Data)
		return false
	}
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
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
