package glfwscreen

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
	"gortex/internal/utils"
)

func (s *GLScreen) RasterShape(shape drawable.Shape) {
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

	for i := 0; i < s.W; i++ {
		for j := 0; j < s.H; j++ {

			// Now (px,py) are coordinates in SCREEN SPACE.
			// We transform them to WORLD through the inverse of the full matrix.
			px := float64(i)/float64(s.W)*2 - 1
			py := float64(j)/float64(s.H)*2 - 1
			px *= float64(s.aspect)

			world := geom.Vector2{X: px, Y: py}
			local := inv.TransformPoint(world)

			if shape.Contains(local.X, local.Y) {
				idx := (i + j*s.W) * 4
				s.buffer[idx+0] = 255 // R
				s.buffer[idx+1] = 0   // G
				s.buffer[idx+2] = 255 // B
				s.buffer[idx+3] = 255 // A
			}
		}
	}
}

func (s *GLScreen) DrawLine(x0, y0, x1, y1 int, color Color) {
	utils.LineBresenham(x0, y0, x1, y1, func(x, y int) { s.SetPixel(x, y, color) })
}
