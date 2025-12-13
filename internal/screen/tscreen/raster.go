package tscreen

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
)

func (s *TermScreen) RasterShape(shape drawable.Shape) {
	// Model matrix describes the position of the shape in WORLD SPACE.
	// For example, a circle in the center or a square rotated at an angle.
	model := shape.ModelMatrix()

	// View matrix describes where the camera is located.
	view := s.Cam.ViewMatrix()

	// Viewport matrix compensates for ASCII pixel distortion.
	// viewport := camera.ViewportMatrix(s.aspect, s.pixelAspect)

	// perform sampling:
	// Full transformation world → screen:
	// final = viewport * view * model
	mv := view.Mul(&model)
	// full := viewport.Mul(&mv)

	// PixelPoint → (inverse transforms) → LocalPoint
	// and then check: shape.Contains(localPoint).
	inv, _ := mv.Inverse()

	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Height; j++ {

			// Now (px,py) are coordinates in SCREEN SPACE.
			// We transform them to WORLD through the inverse of the full matrix.
			px := float64(i)/float64(s.Width)*2 - 1
			py := float64(j)/float64(s.Height)*2 - 1
			px *= s.aspect * s.pixelAspect

			world := geom.Vector2{X: px, Y: py}
			local := inv.TransformPoint(world)

			if shape.Contains(local.X, local.Y) {
				s.buffer[i+j*s.Width] = '@'
			}
		}
	}
}
