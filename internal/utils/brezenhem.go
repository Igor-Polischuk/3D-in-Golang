package utils

type SetPixel func(x, y int)

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
