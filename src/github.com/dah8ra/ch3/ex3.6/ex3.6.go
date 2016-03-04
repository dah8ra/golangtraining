package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("lena.jpg")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// resize for performance
	//	img = resize.Resize(36, 0, img, resize.Lanczos3)

	width, height := getImageDimension("lena.jpg")
	fmt.Println("Width:", width, "Height:", height)

	bsize := 4
	w := width / bsize
	h := height / bsize

	outimg := getOutimg(w, h, bsize, img)

	var outFile *os.File
	if outFile, err = os.Create("out.jpg"); err != nil {
		println("Error", err)
		return
	}

	option := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(outFile, outimg, option); err != nil {
		//if err = jpeg.Encode(outFile, img, nil); err != nil {
		println()
		return
	}

	defer outFile.Close()
}

func getOutimg(width int, height int, bsize int, img image.Image) image.Image {
	rect := image.Rect(0, 0, width, height)
	outimg := image.NewRGBA(rect)
	for x := 0; x < width; x += bsize {
		for y := 0; y < height; y += bsize {
			var top color.Color
			var bottom color.Color
			var left color.Color
			var right color.Color
			if x > 1 && y > 1 && x < width && y < height {
				top = img.At(x, y-1)
				bottom = img.At(x, y+1)
				//				center := img.At(x, y)
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == 0 && y == 0 { // hidariue
				//top := img.At(x,y-1)
				bottom = img.At(x, y+1)
				//				center := img.At(x, y)
				//left := img.At(x-1,y)
				right = img.At(x+1, y)
			} else if x > 1 && x < width && y == 0 { // ue
				//top := img.At(x, y-1)
				bottom = img.At(x, y+1)
				//				center := img.At(x, y)
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == width && y == 0 { // migiue
				//top := img.At(x, y-1)
				bottom = img.At(x, y+1)
				//				center := img.At(x, y)
				left = img.At(x-1, y)
				//right := img.At(x+1, y)
			} else if x == 0 && y > 1 && y < height { //hidari
				top = img.At(x, y-1)
				bottom = img.At(x, y+1)
				//				center := img.At(x, y)
				//left := img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == 0 && y == height { // hidarishita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
				//				center := img.At(x, y)
				//left := img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x > 1 && y == height && x < width { // shita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
				//				center = img.At(x, y)
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == width && y == height { // migishita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
				//				center = img.At(x, y)
				left = img.At(x-1, y)
				//right := img.At(x+1, y)
			}
			top = getNilToGrayRGBA(top)
			bottom = getNilToGrayRGBA(bottom)
			left = getNilToGrayRGBA(left)
			right = getNilToGrayRGBA(right)
			center := getSamplingColor(top, bottom, left, right)
			for i := 0; i < bsize; i++ {
				for j := 0; j < bsize; j++ {
					outimg.Set(x+i, y+j, center)
				}
			}
		}
	}
	return outimg
}

func getNilToGrayRGBA(c color.Color) color.Color {
	if c == nil {
		return color.RGBA{R: 128, G: 128, B: 128, A: 128}
	}
	return c
}

//	max := 0
// find most appered color
//	for _, val := range hist {
//		if val > max {
//			max = val
//		}
//	}
//	r, g, b := int2rgb(max)
//	colorCode := "#" + dec2hex(r, 2) + dec2hex(g, 2) + dec2hex(b, 2)
//
//	 most appered color in image
//	println(colorCode)

func getSamplingColor(top color.Color, bottom color.Color, left color.Color, right color.Color) color.Color {
	tr, tg, tb, ta := top.RGBA()
	br, bg, bb, ba := bottom.RGBA()
	lr, lg, lb, la := left.RGBA()
	rr, rg, rb, ra := right.RGBA()
	cr := (tr + br + lr + rr) / 4
	cg := (tg + bg + lg + rg) / 4
	cb := (tb + bb + lb + rb) / 4
	ca := (ta + ba + la + ra) / 4
	ucr := uint8(cr)
	ucg := uint8(cg)
	ucb := uint8(cb)
	uca := uint8(ca)
	return color.RGBA{R: ucr, G: ucg, B: ucb, A: uca}
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func getHistgram(img image.Image) []int {
	hist := make([]int, 1000000)

	// get bounds
	rect := img.Bounds()
	// color reduction
	for i := 0; i < rect.Max.Y; i++ {
		for j := 0; j < rect.Max.X; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			i := rgb2int(int(r), int(g), int(b))
			hist[i]++
		}
	}
	return hist
}

func dec2hex(n, beam int) string {
	hex := ""
	str := "0123456789abcdef"
	for i := 0; i < beam; i++ {
		m := n & 0xf
		hex = string(str[m]) + hex
		n -= m
		n >>= 4
	}
	return hex
}

func rgb2int(r, g, b int) int {
	return (((r >> 5) << 6) | ((g >> 5) << 3) | ((b >> 5) << 0))
}

func int2rgb(i int) (r, g, b int) {
	return ((i >> 6 & 0x7) << 5) + 16, ((i >> 3 & 0x7) << 5) + 16, ((i >> 0 & 0x7) << 5) + 16
}
