package shapes

import (
	"gortex/internal/drawable"
	"gortex/internal/geom"
)

type Rectangle struct {
	Pos   geom.Vector2
	Size  geom.Vector2
	Angle int16
}

func (r Rectangle) Rasterize(ctx drawable.RenderContext) {
	ctx.Rasterizer.RasterRect(r.Pos.X, r.Pos.Y, r.Size.X, r.Size.Y, r.Angle, ctx.Texture)
}
