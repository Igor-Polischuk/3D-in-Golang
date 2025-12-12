package drawable

type Rasterizer interface {
	RasterCircle(x, y, r float64, texture []rune)
	RasterRect(x, y, w, h float64, angle int16, texture []rune)
}

type RenderContext struct {
	Rasterizer Rasterizer
	Texture    []rune
}

type Drawable interface {
	Rasterize(ctx RenderContext)
}
