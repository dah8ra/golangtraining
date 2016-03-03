// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var red color.Color = color.RGBA{255, 128, 128, 128}
var green color.Color = color.RGBA{128, 255, 128, 128}
var blue color.Color = color.RGBA{128, 128, 255, 128}
var palette = []color.Color{color.Black, red, green, blue, color.White}

const (
	blackIndex = 4
	redIndex   = 0
	greenIndex = 2
	blueIndex  = 1
	whiteIndex = 3
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
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
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

func getColor(n uint8) color.Color {
	temp := n % 5
	switch temp {
	case blackIndex:
		return palette[blackIndex]
	case redIndex:
		return palette[redIndex]
	case greenIndex:
		return palette[greenIndex]
	case blueIndex:
		return palette[blueIndex]
	case whiteIndex:
		return palette[whiteIndex]
	}
	return palette[redIndex]
}

//!-

// Some other interesting functions:

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

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
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
