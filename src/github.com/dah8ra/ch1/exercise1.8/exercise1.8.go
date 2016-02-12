package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpStr := "http://"
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, httpStr) == false {
			fmt.Println("Added http prefix!!!")
			url = httpStr + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		w := bufio.NewWriter(os.Stdout)
		_, err = io.Copy(w, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
