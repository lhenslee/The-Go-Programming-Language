package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale = height*0.4
	angle = math.Pi/6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'"+
				"style='stroke: grey; fill: white; stroke-width: 0.7'" +
				"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			color := 0xff0000
			if ay > height/2 {
				color = 0x0000ff
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:#%06x'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange*(float64(i)/cells-0.5)
	y := xyrange*(float64(j)/cells-0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		return math.NaN(), math.NaN()
	}
	
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2+(x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale-z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Cos(r)
}