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
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
