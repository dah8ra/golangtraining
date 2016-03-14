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
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	words, images := countWordsAndImage(doc)
	fmt.Printf("words: %d, images: %d\n", words, images)
}

func countWordsAndImage(n *html.Node) (words, images int) {
	if strings.Contains(n.Data, "img") {
		images++
	}
	if n.Type == html.TextNode {
		words += countCharacter(n.Data)
		//		fmt.Printf("%s", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retWords, retImages := countWordsAndImage(c)
		words += retWords
		images += retImages
	}
	return words, images
}

func countCharacter(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(string(line)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
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
