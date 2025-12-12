package drawable

type Rasterizer interface {
	RasterCircle(x, y, r float64, texture []rune)
}

type RenderContext struct {
	Rasterizer Rasterizer
	Texture    []rune
}

type Drawable interface {
	Rasterize(ctx RenderContext)
}
