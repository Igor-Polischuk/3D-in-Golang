package drawable

type Rasterizer interface {
	RasterCircle(x, y, r float32)
}

type Drawable interface {
	Rasterize(r Rasterizer)
}
