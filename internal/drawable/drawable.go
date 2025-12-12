package drawable

import "gortex/internal/geom"

type ShapeModel = geom.Matrix
type ShapeContainsFunction = func(x, y float64) bool

type Shape interface {
	Contains(x, y float64) bool
	ModelMatrix() geom.Matrix
}

type Rasterizer interface {
	Raster(shape Shape)
}

type RenderContext struct {
	Rasterizer Rasterizer
	Texture    []rune
}

type Drawable interface {
	Rasterize(ctx RenderContext)
}
