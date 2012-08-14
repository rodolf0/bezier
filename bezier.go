package main

type Point struct {
	x, y float32
}

type Curve func(float32) Point

func interpolate(p0, p1 Point, t float32) Point {
	return Point{p0.x + t*(p1.x-p0.x), p0.y + t*(p1.y-p0.y)}
}

// return a parametric bezier function constructed from pts
func NewCurve(pts ...Point) Curve {
	if len(pts) < 2 {
		return nil
	} else if len(pts) == 2 {
		return func(t float32) Point {
			return interpolate(pts[0], pts[1], t)
		}
	} else if len(pts) == 3 {
		return func(t float32) Point {
			var a, b = interpolate(pts[0], pts[1], t), interpolate(pts[1], pts[2], t)
			return interpolate(a, b, t)
		}
	}
	var mid_pts = append(pts, make([]Point, len(pts)*(len(pts)-1)/2)...)
	return func(t float32) Point {
		for m, n := len(pts), len(pts)-1; n > 0; n-- {
			for i := 0; i < n; i++ {
				mid_pts[m+i] = interpolate(mid_pts[m+i-n-1], mid_pts[m+i-n], t)
			}
			m += n
		}
		return mid_pts[len(mid_pts)-1]
	}
}
