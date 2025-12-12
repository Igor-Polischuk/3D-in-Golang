package screen

import "math"

func (s *TermScreen) RasterCircle(x, y, r float64, texture []rune) {
	for i := range s.Width {
		for j := range s.Height {
			dx := float64(i)/float64(s.Width)*2 - 1
			dy := float64(j)/float64(s.Height)*2 - 1
			dx -= x
			dy -= y

			dx *= float64(s.aspect) * float64(s.pixelAspect)

			dist := math.Sqrt((float64(dx*dx + dy*dy)))
			color := int(1 / dist)

			if color <= 0 {
				color = 0
			} else if color >= len(texture) {
				color = len(texture) - 1
			}

			pixel := texture[color]

			if dx*dx+dy*dy < r {
				s.buffer[i+j*s.Width] = pixel
			}
		}
	}
}
