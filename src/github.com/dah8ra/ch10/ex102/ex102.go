package main

import (
	_ "archive/tar"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	//	readTar()
	readZip()
}

/*
func readTar() {
	var file *os.File
	var err error

	if file, err = os.Open("test.tar"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := tar.NewReader(file)

	var header *tar.Header
	for {
		header, err = reader.Next()
		if err == io.EOF {
			// ファイルの最後
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		s := header.Name
		fmt.Printf("@@@ %s\n", s)
	}
}
*/

type Decompressor func(r io.Reader) io.ReadCloser

func init() {
	zip.RegisterDecompressor(0, Decompressor(zip.Reader))
}

func readZip() {
	reader, err := zip.OpenReader("test.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	var rc io.ReadCloser
	for _, f := range reader.File {
		rc, err = f.Open()
		if err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		if err != nil {
			log.Fatal(err)
		}

		s := f.Name
		fmt.Printf("@@@ %s\n", s)
		rc.Close()
	}
}
