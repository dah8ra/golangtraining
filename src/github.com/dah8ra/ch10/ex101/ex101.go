// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

var s = flag.String("type", "jpeg", "Select output image type. PNF, JPEG suported.")

// $go run ex101.go -type jpeg < test.png > out.jpeg
func main() {
	flag.Parse()

	if err := convert(os.Stdin, os.Stdout, *s); err != nil {
		fmt.Sprintln("Type: %s\n", *s)
		fmt.Fprintf(os.Stderr, "Specified output type: %s, err: %v\n", *s, err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, s string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
