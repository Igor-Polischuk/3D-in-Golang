package shapes

import (
	"gortex/internal/geom"
)

type Circle struct {
	Pos geom.Vector2
	R   float64
}

func (c Circle) Contains(x, y float64) bool {
	return x*x+y*y <= 1
}

func (c Circle) ModelMatrix() geom.Matrix {
	translate := geom.Translate(c.Pos.X, c.Pos.Y)
	scale := geom.Scale(c.R, c.R)

	return translate.Mul(&scale)
}
