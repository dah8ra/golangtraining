package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/"  || local == "."{
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, writeerr := write(local, f, resp)
	return local, n, writeerr
}

func write(local string, f *os.File, resp *http.Response) (int64, error) {
	n, err := io.Copy(f, resp.Body)
	defer func() {
		if err == nil {
			fmt.Println("Copy error is nil.")
			err = f.Close()
			if err != nil {
				fmt.Println("Input file close error.")
			}
		}
	}()
	return n, err
}
