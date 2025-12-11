package screen

func (s *TermScreen) RasterCircle(x, y, r float32) {
	for i := range s.width {
		for j := range s.height {
			dx := float32(i)/float32(s.width)*2 - 1
			dy := float32(j)/float32(s.height)*2 - 1
			dx -= x
			dy -= y

			dx *= float32(s.aspect) * float32(s.pixelAspect)

			if dx*dx+dy*dy < r {
				s.buffer[i+j*s.width] = '@'
			}
		}
	}
}
