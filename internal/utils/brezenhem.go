package utils

import "math"

type SetPixel func(x, y int)
type SetPixelT func(x, y int)

func LineBresenham(x0, y0, x1, y1 int, setPixel SetPixel) {
	dx := x1 - x0
	dy := y1 - y0

	sx := 1
	if dx < 0 {
		sx = -1
		dx = -dx
	}
	sy := 1
	if dy < 0 {
		sy = -1
		dy = -dy
	}

	err := dx - dy

	x, y := x0, y0

	for {
		setPixel(x, y)
		if x == x1 && y == y1 {
			break
		}

		e2 := 2 * err

		if e2 > -dy {
			err -= dy
			x += sx
		}

		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

func BresenhamWithT(x0, y0, x1, y1 int, plot func(x, y int, t float64)) {
	dx := math.Abs(float64(x1 - x0))
	dy := math.Abs(float64(y1 - y0))
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}

	err := dx - dy
	length := math.Max(dx, dy)
	step := 0.0

	for {
		t := step / length
		plot(x0, y0, t)

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}

		step++
	}
}
