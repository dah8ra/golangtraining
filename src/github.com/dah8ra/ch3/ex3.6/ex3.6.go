package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	const (
		blocksize = 4
	)

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

	width, height := getImageDimension("lena.jpg")
	//fmt.Println("Width:", width, "Height:", height)

	outimg := supersampling(width, height, blocksize, img)
	//	outimg := getOutimg(width, height, blocksize, img)

	var outFile *os.File
	if outFile, err = os.Create("out.jpg"); err != nil {
		println("Error", err)
		return
	}

	option := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(outFile, outimg, option); err != nil {
		return
	}

	defer outFile.Close()
}

func supersampling(width int, height int, bsize int, img image.Image) image.Image {
	rect := image.Rect(0, 0, width*bsize, height*bsize)
	outimg := image.NewRGBA(rect)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := img.At(x, y)
			for i := 0; i < bsize; i++ {
				for j := 0; j < bsize; j++ {
					outimg.Set(x*bsize+i, y*bsize+j, c)
				}
			}
		}
	}
	return outimg
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
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == 0 && y == 0 { // hidariue
				//top := img.At(x,y-1)
				bottom = img.At(x, y+1)
				//left := img.At(x-1,y)
				right = img.At(x+1, y)
			} else if x > 1 && x < width && y == 0 { // ue
				//top := img.At(x, y-1)
				bottom = img.At(x, y+1)
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == width && y == 0 { // migiue
				//top := img.At(x, y-1)
				bottom = img.At(x, y+1)
				left = img.At(x-1, y)
				//right := img.At(x+1, y)
			} else if x == 0 && y > 1 && y < height { //hidari
				top = img.At(x, y-1)
				bottom = img.At(x, y+1)
				//left := img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == 0 && y == height { // hidarishita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
				//left := img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x > 1 && y == height && x < width { // shita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
				left = img.At(x-1, y)
				right = img.At(x+1, y)
			} else if x == width && y == height { // migishita
				top = img.At(x, y-1)
				//bottom := img.At(x, y+1)
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
