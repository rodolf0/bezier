package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// draw a bezier curve with it's control points
func drawCurve(img draw.Image, c color.Color, pts ...Point) {
	const segments = 50
	var f = NewCurve(pts...)

	for prev, t := f(0.0), 1; t <= segments; t++ {
		var cur = f(float32(t) / float32(segments))
		line(int(cur.x), int(cur.y), int(prev.x), int(prev.y), c, img)
		prev = cur
	}
}

// connect points with lines
func drawPolyline(img draw.Image, c color.Color, pts ...Point) {
	for j := 1; j < len(pts); j++ {
		line(int(pts[j-1].x), int(pts[j-1].y), int(pts[j].x), int(pts[j].y), c, img)
	}
}

func whiteImage() *image.RGBA {
	var i = image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(i, i.Rect, image.NewUniform(color.RGBA{255, 255, 255, 255}), image.Point{0, 0}, draw.Src)
	return i
}

func main() {
	// http://jeremykun.com/2013/05/11/bezier-curves-and-picasso/
	var pts = [][]Point{
		{{180, 280}, {183, 268}, {186, 256}, {189, 244}}, // front leg
		{{191, 244}, {290, 244}, {300, 230}, {339, 245}}, // tummy
		{{340, 246}, {350, 290}, {360, 300}, {355, 210}}, // back leg
		{{353, 210}, {370, 207}, {380, 196}, {375, 193}}, // tail
		{{375, 193}, {310, 220}, {190, 220}, {164, 205}}, // back
		{{164, 205}, {135, 194}, {135, 265}, {153, 275}}, // ear start
		{{153, 275}, {168, 275}, {170, 180}, {150, 190}}, // ear end + head
		{{149, 190}, {122, 214}, {142, 204}, {85, 240}},  // nose bridge
		{{86, 240}, {100, 247}, {125, 233}, {140, 238}},  // mouth
	}

	var img = whiteImage()
	for _, curve := range pts {
		drawCurve(img, color.RGBA{40, 40, 50, 255}, curve...)
		/*drawPolyline(img, color.RGBA{230, 230, 230, 255}, curve...)*/
	}

	out, _ := os.Create("__test-out.png")
	png.Encode(out, img)
}
