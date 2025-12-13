package tscreen

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
	"gortex/internal/utils"
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

	for i := 0; i < s.W; i++ {
		for j := 0; j < s.H; j++ {

			// Now (px,py) are coordinates in SCREEN SPACE.
			// We transform them to WORLD through the inverse of the full matrix.
			px := float64(i)/float64(s.W)*2 - 1
			py := 1 - float64(j)/float64(s.H)*2

			px *= s.aspect * s.pixelAspect

			world := geom.Vector2{X: px, Y: py}
			local := inv.TransformPoint(world)

			if shape.Contains(local.X, local.Y) {
				s.buffer[i+j*s.W] = '@'
			}
		}
	}
}

func (s TermScreen) DrawLine(x0, y0, x1, y1 int) {
	y0_raster := s.H - 1 - y0
	y1_raster := s.H - 1 - y1
	utils.LineBresenham(x0, y0_raster, x1, y1_raster, func(x, y int) { s.SetPixel(x, y, '@') })
}
