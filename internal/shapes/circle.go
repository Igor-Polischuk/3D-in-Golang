package shapes

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
)

type Circle struct {
	Pos geom.Vector2
	R   float64
}

func (c *Circle) IsPointWithin(x, y float64) bool {
	return x*x+y*y < c.R
}

func (c Circle) Rasterize(ctx drawable.RenderContext) {
	// r := ctx.Rasterizer
	// r.RasterCircle(float32(c.Pos.X), float32(c.Pos.Y), c.R)
	ctx.Rasterizer.RasterCircle(c.Pos.X, c.Pos.Y, c.R, ctx.Texture)
}
