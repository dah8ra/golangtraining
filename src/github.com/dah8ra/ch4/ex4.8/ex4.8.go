// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letterCounts := make(map[rune]int)    // counts of Unicode characters
	markCounts := make(map[rune]int)
	numberCounts := make(map[rune]int)
	digitCounts := make(map[rune]int)
	graphicCounts := make(map[rune]int)
	spaceCounts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int       // count of lengths of UTF-8 encodings
	invalid := 0                          // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error		
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		} else if unicode.IsMark(r){
			markCounts[r]++
		} else if unicode.IsNumber(r){
			numberCounts[r]++
		} else if unicode.IsDigit(r) {
			digitCounts[r]++
		} else if unicode.IsGraphic(r){
			graphicCounts[r]++
		} else if unicode.IsSpace(r) {
			spaceCounts[r]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range letterCounts {
		fmt.Printf("letter: %q\t%d\n", c, n)
	}
	for c, n := range markCounts {
		fmt.Printf("mark: %q\t%d\n", c, n)
	}
	for c, n := range numberCounts {
		fmt.Printf("number: %q\t%d\n", c, n)
	}
	for c, n := range digitCounts {
		fmt.Printf("digit: %q\t%d\n", c, n)
	}
	for c, n := range graphicCounts {
		fmt.Printf("graphic: %q\t%d\n", c, n)
	}
	for c, n := range spaceCounts {
		fmt.Printf("space: %q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
