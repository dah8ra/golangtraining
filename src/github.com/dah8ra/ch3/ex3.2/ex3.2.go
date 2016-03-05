package main

import (
	"fmt"
	"math"
)

const (
	//width, height = 300, 160                // canvas size in pixels
	//cells   = 50                      // number of grid cells
	//xyrange = 15.0                    // axis ranges (-xyrange..+xyrange)
	//xyscale = width / 2 / xyrange * 5 // pixels per x or y unit

	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if ax == -1 || bx == -1 || cx == -1 || dx == -1 {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := saddle(x, y)
	//		z, ok := eggbox(x, y)
	//	z, ok := mogle(x, y)
	//	z, ok := gaussian(x, y)
	if !ok {
		return -1, -1
	}
	//z := xyscale / 100 * zz

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func saddle(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := (math.Pow(x, 2) - math.Pow(y, 2)) / (20 * r)
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func eggbox(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	rad := math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) / (20 * r)
	s := rad * math.Sin(x)
	c := rad * math.Cos(y)
	result := math.Sqrt(math.Pow(s, 2) + math.Pow(c, 2))
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func mogle(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	rad := math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) / (20 * r)
	result := rad * math.Sin(x)
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func gaussian(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Exp(-(r * r))
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func wave(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(x+y) / r
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func sqrtmapping(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sqrt(1+math.Pow(x, 2)+math.Pow(y, 2)) / (2 * r)
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func eggbox2(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	rad := math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) / (2 * r)

	s := rad * math.Sin(x)
	c := rad * math.Cos(y)
	t := math.Sqrt(math.Pow(s, 2) + math.Pow(c, 2))

	result := rad * math.Cos(t)

	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}
