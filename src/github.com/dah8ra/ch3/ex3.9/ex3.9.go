package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fractal(w, "0", "0", "1")
	})
	http.HandleFunc("/fractal", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.RawQuery
		splitEqual := strings.Split(query, "=")
		splitComma := strings.Split(splitEqual[1], ",")
		fractal(w, splitComma[0], splitComma[1], splitComma[2])
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func fractal(out io.Writer, xstr string, ystr string, magstr string) {
	xxfloat, _ := strconv.ParseFloat(xstr, 64)
	yyfloat, _ := strconv.ParseFloat(ystr, 64)
	magfloat, _ := strconv.ParseFloat(magstr, 64)
	xx := int(xxfloat)
	yy := int(yyfloat)
	mag := int(magfloat)
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width*mag, height*mag))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px*mag, py*mag, newton(z))
		}
	}
	outimg := image.NewRGBA(image.Rect(0, 0, width, height))
	cx := width*mag/2 + xx
	cy := height*mag/2 + yy
	ccx := cx - width/2
	ccy := cy - height/2
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			outimg.Set(px, py, img.At(ccx+px, ccy+py))
		}
	}
	png.Encode(out, outimg)
	fmt.Printf("Done!!! (x,y)=(%d,%d), magnification: %d", xx, yy, mag)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		//		fmt.Printf("i: %d, z: %5.5f\n", i, z)
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// return color.Gray{255 - contrast*i}
			r := contrast * i * 2
			g := contrast * i * 10
			b := contrast * i * 8
			a := contrast * i * 10
			return color.RGBA{r, g, b, a}
		}
	}
	return color.Black
}
