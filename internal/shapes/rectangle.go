package shapes

import (
	"gortex/internal/geom"
	"math"
)

type Rectangle struct {
	Pos   geom.Vector2
	Size  geom.Vector2
	Angle int16
}

func (r Rectangle) ModelMatrix() geom.Matrix {
	translate := geom.Translate(r.Pos.X, r.Pos.Y)
	rotate := geom.Rotate(float64(r.Angle))
	scale := geom.Scale(r.Size.X, r.Size.Y)

	return translate.Mul(&rotate).Mul(&scale)
}

func (r Rectangle) Contains(x, y float64) bool {
	return math.Abs(x) <= 0.5 && math.Abs(y) <= 0.5
}
