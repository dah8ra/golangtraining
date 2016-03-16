package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	reader := getFileContents()
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	displayAttribute(n)

	if post != nil {
		post(n)
	}
}

func displayAttribute(n *html.Node) {
	for _, a := range n.Attr {
		fmt.Printf("%*s%s\n", depth*2, "", a.Val)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if len(n.Attr) != 0 {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
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
