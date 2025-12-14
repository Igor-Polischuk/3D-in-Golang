package render

import (
	"gortex/internal/geom"
	"gortex/internal/mesh"
	"gortex/internal/scene"
)

type Screen interface {
	DrawLine(x0, y0, x1, y1 int)
	Width() int
	Height() int
}

type Renderer struct {
	depth  []float64
	screen Screen
}

func NewRenderer(screen Screen) Renderer {
	return Renderer{
		depth:  []float64{},
		screen: screen,
	}
}

func (r Renderer) Render(scn scene.Scene, vp geom.Matrix) {
	for _, entity := range scn.Entities {
		model := entity.ModelMatrix()
		MVP := vp.Mul(&model)

		r.RenderMesh(entity.Mesh, MVP)
	}
}

func ToScreen(v geom.Vector3, w, h int) (int, int) {
	x := (v.X + 1) * 0.5 * float64(w-1)
	y := (v.Y + 1) * 0.5 * float64(h-1)

	return int(x), int(y)
}

func (r Renderer) RenderMesh(mesh mesh.Mesh, MVP geom.Matrix) {
	// MVP := proj.Mul(&view).Mul(&model)
	transformed := mesh.Transform(MVP)

	for i := 0; i < len(mesh.Indices); i += 3 {
		a := transformed[mesh.Indices[i]]
		b := transformed[mesh.Indices[i+1]]
		c := transformed[mesh.Indices[i+2]]

		// project to screen
		ax, ay := ToScreen(a, r.screen.Width(), r.screen.Height())
		bx, by := ToScreen(b, r.screen.Width(), r.screen.Height())
		cx, cy := ToScreen(c, r.screen.Width(), r.screen.Height())

		// draw wireframe
		r.screen.DrawLine(ax, ay, bx, by)
		r.screen.DrawLine(bx, by, cx, cy)
		r.screen.DrawLine(cx, cy, ax, ay)
	}
}
