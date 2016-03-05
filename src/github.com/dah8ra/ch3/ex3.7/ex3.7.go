package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		mag                    = 1
		xx                     = 0
		yy                     = 0
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
	png.Encode(os.Stdout, outimg)
	//	png.Encode(os.Stdout, img)
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

func mynewton() {
	var z1 float64 = 5
	var z2 float64 = 0
	var i int = 0
	for {
		z2 = z1 - (math.Pow(z1, 4)-1)/(4*math.Pow(z1, 3))
		z := z2 - z1
		fmt.Printf("i: %d, x: %5.5f\n", i, math.Abs(z))
		if math.Abs(z) < 0.0001 {
			break
		}
		z1 = z2
		i++
	}

	fmt.Print(z2)
}
