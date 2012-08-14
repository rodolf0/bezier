package main

import (
	"os"
	"image"
	"image/png"
	"image/color"
	"image/draw"
)

// draw a bezier curve with it's control points
func drawCurve(img draw.Image, c color.Color, pts ...Point) {
	const segments = 50
	var f = NewCurve(pts...)

	for prev, t := f(0.0), 1; t <= segments; t++ {
		var cur = f(float32(t)/float32(segments))
		line(int(cur.x), img.Bounds().Max.Y - 1 - int(cur.y),
				 int(prev.x), img.Bounds().Max.Y - 1 - int(prev.y), c, img)
		prev = cur
	}
}

// connect points with lines
func drawPolyline(img draw.Image, c color.Color, pts ...Point) {
	for j := 1; j < len(pts); j++ {
		line(int(pts[j-1].x), img.Bounds().Max.Y - 1 - int(pts[j-1].y),
				 int(pts[j].x), img.Bounds().Max.Y - 1 - int(pts[j].y), c, img)
	}
}

func whiteImage() *image.RGBA {
	var i = image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(i, i.Rect, image.NewUniform(color.RGBA{255, 255, 255, 255}), image.Point{0, 0}, draw.Src)
	return i
}

func main() {
	var pts = []Point{{0, 0}, {499, 100}, {0, 100}, {0, 400}, {499, 499}}
	var pts2 = []Point{{0, 10}, {150, 500}, {300, 30}, {450, 500}}
	var pts3 = []Point{{0, 0}, {256, 499}, {499, 0}}
	var pts4 = []Point{{0, 499}, {499, 255}}
	var pts5 = []Point{{0, 0}, {1000, 499}, {-500, 499}, {499, 0}}

	var img = whiteImage()
	drawCurve(img, color.RGBA{255, 200, 50, 255}, pts...)
	drawPolyline(img, color.RGBA{30, 255, 30, 255}, pts...)

	drawCurve(img, color.RGBA{50, 50, 255, 255}, pts2...)
	drawPolyline(img, color.RGBA{255, 0, 0, 255}, pts2...)

	drawCurve(img, color.RGBA{50, 50, 50, 255}, pts3...)
	drawPolyline(img, color.RGBA{128, 128, 128, 255}, pts3...)

	drawCurve(img, color.RGBA{50, 250, 250, 255}, pts4...)
	drawPolyline(img, color.RGBA{50, 50, 100, 255}, pts4...)

	drawCurve(img, color.RGBA{140, 200, 78, 255}, pts5...)
	drawPolyline(img, color.RGBA{89, 165, 100, 255}, pts5...)

	out, _ := os.Create("__test-out.png")
	png.Encode(out, img)
}
