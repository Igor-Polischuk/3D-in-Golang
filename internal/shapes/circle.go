package shapes

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
)

type Circle struct {
	Pos geom.Vector2
	R   float32
}

func (c *Circle) IsPointWithin(x, y float32) bool {
	return x*x+y*y < c.R
}

func (c Circle) Rasterize(r drawable.Rasterizer) {
	r.RasterCircle(float32(c.Pos.X), float32(c.Pos.Y), c.R)
}
