package screen

func (s *TermScreen) RasterCircle(x, y, r float32) {
	for i := range s.width {
		for j := range s.height {
			x := float32(i)/float32(s.width)*2 - 1
			y := float32(j)/float32(s.height)*2 - 1
			x *= float32(s.aspect) * float32(s.pixelAspect)

			if x*x+y*y < 0.5 {
				s.buffer[i+j*s.width] = '@'
			}
		}
	}
}
