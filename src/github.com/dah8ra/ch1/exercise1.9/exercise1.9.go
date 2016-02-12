package main

import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(resp.Body)
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("ResponseCode: %s\n============================\n%s", status, b)

	}
}
