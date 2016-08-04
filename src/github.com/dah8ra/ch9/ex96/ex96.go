package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

///////////////////////////////////////////
// Created 100 gif files.
// Default     : 13.6407802 seconds
// GOMAXPROCS=1: 12.3477062 seconds
// GOMAXPROCS=2: 13.0807482 seconds
// GOMAXPROCS=3: 14.0348028 seconds
// GOMAXPROCS=4: 13.0327455 seconds
///////////////////////////////////////////
func main() {
	elapsed := time.Now()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		name := "test_" + strconv.Itoa(i) + ".gif"
		output, _ := os.Create(name)
		writer := bufio.NewWriter(output)
		go start(writer)
	}
	fmt.Println("Receiving...")
	wg.Wait()
	fmt.Printf("Elapsed %s seconds...\n", time.Since(elapsed))
	fmt.Println("Done")
}

func start(writer io.Writer) {
	defer wg.Done()
	const (
		xmin, ymin, xmax, ymax = -4, -4, +4, +4
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(writer, img) // NOTE: ignoring errors	
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint64(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			var temp uint8 = uint8(50 * n)
			return color.RGBA{0, temp, temp, 255}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
