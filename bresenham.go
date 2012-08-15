package main

import (
	"image/color"
	"image/draw"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// http://en.wikipedia.org/wiki/Bresenham's_line_algorithm#Simplification
func line(x0, y0, x1, y1 int, c color.Color, img draw.Image) {
	var dx = abs(x1 - x0)
	var dy = abs(y1 - y0)
	var err = dx - dy
	var sx, sy = 1, 1

	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}

	img.Set(x0, y0, c)
	for x0 != x1 || y0 != y1 {
		var e2 = 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
		img.Set(x0, y0, c)
	}

}
