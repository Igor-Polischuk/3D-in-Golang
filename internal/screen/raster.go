package screen

import (
	"math"
)

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

func (s *TermScreen) RasterRect(x, y, w, h float64, angleDeg int16, texture []rune) {
	border := texture[0]
	fill := texture[1]

	// convert angle to radians
	angle := float64(angleDeg) * math.Pi / 180

	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// border thickness
	epsX := 2.0 / float64(s.Width)
	epsY := 2.0 / float64(s.Height)

	// center of rectangle
	cx := x + w/2
	cy := y + h/2

	halfW := w / 2
	halfH := h / 2

	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Height; j++ {

			// normalize pixel
			px := float64(i)/float64(s.Width)*2 - 1
			py := float64(j)/float64(s.Height)*2 - 1

			// aspect correction
			px *= float64(s.aspect) * float64(s.pixelAspect)

			// translate into rectangle-local space
			dx := px - cx
			dy := py - cy

			// inverse rotation
			lx := dx*cosA + dy*sinA
			ly := -dx*sinA + dy*cosA

			// inside rectangle?
			if lx >= -halfW && lx <= halfW && ly >= -halfH && ly <= halfH {

				// detect border (still in local coords!)
				onBorder :=
					math.Abs(lx+halfW) < epsX ||
						math.Abs(lx-halfW) < epsX ||
						math.Abs(ly+halfH) < epsY ||
						math.Abs(ly-halfH) < epsY

				if onBorder {
					s.buffer[i+j*s.Width] = border
				} else {
					s.buffer[i+j*s.Width] = fill
				}
			}
		}
	}
}
