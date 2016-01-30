package main 

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var red color.Color = color.RGBA{255, 128, 128, 128}
var green color.Color = color.RGBA{128, 255, 128, 128}
var blue color.Color = color.RGBA{128, 128, 255, 128}
var palette = []color.Color{color.Black, red, green, blue, color.White}

const (
	blackIndex = 0
	redIndex = 1
	greenIndex = 2
	blueIndex = 3
	whiteIndex = 4
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)	
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex byte
		if i % 3 == 0 {
			colorIndex = redIndex
		} else if i % 3 == 1 {
			colorIndex = greenIndex
		} else if i % 3 == 2{
			colorIndex = blueIndex
		} else {
			colorIndex = whiteIndex
		}
		
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)		
	}
	gif.EncodeAll(out, &anim)
}