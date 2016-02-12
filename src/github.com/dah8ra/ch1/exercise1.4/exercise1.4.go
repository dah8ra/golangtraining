package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for i := 0; i < len(files); i++ {
			arg := files[i]
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}

}

func countLines(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("file: %s, %d\t%s\n", f.Name(), n, line)
		}
	}
}
