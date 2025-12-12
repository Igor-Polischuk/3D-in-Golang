package screen

import "math"

func (s *TermScreen) RasterCircle(x, y, r float32) {
	tex := s.currentTexture

	for i := range s.width {
		for j := range s.height {
			dx := float32(i)/float32(s.width)*2 - 1
			dy := float32(j)/float32(s.height)*2 - 1
			dx -= x
			dy -= y

			dx *= float32(s.aspect) * float32(s.pixelAspect)

			dist := math.Sqrt((float64(dx*dx + dy*dy)))
			color := int(1 / dist)

			if color <= 0 {
				color = 0
			} else if color >= len(tex) {
				color = len(tex) - 1
			}

			pixel := tex[color]

			if dx*dx+dy*dy < r {
				s.buffer[i+j*s.width] = pixel
			}
		}
	}
}
