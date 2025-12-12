package screen

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
)

func (s *TermScreen) RasterShape(shape drawable.Shape) {
	model := shape.ModelMatrix()

	inv, _ := model.Inverse()

	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Height; j++ {

			// 1. normalize pixel into -1..1
			px := float64(i)/float64(s.Width)*2 - 1
			py := float64(j)/float64(s.Height)*2 - 1
			px *= s.aspect * s.pixelAspect
			// 2. APPLY CAMERA HERE LATER (view matrix)
			// px, py = view.Transform(px, py)

			// 3. localPoint = inverseModel * screenPoint
			local := inv.TransformPoint(geom.Vector2{X: px, Y: py})

			if shape.Contains(local.X, local.Y) {
				s.buffer[i+j*s.Width] = '@'
			}
		}
	}
}
