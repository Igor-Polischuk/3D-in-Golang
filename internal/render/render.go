package render

import (
	"gortex/internal/geom"
	"gortex/internal/material"
	"gortex/internal/mesh"
	"gortex/internal/scene"
	"gortex/internal/utils"
	"math"
)

type Material interface {
	Shade(input material.FragmentInput) material.Pixel
}

type Screen interface {
	Width() int
	Height() int
	SetPixel(x, y int, pixel material.Pixel)
}

type Renderer struct {
	depth  []float64
	screen Screen
}

func NewRenderer(screen Screen) Renderer {
	return Renderer{
		depth:  make([]float64, screen.Width()*screen.Height()),
		screen: screen,
	}
}

func (r Renderer) Render(scn scene.Scene, vp geom.Matrix) {
	for i := range r.depth {
		r.depth[i] = math.Inf(1) // +∞
	}
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

func (r Renderer) setPixel(x, y int, z float64) {
	idx := x + y*r.screen.Width()
	if z < r.depth[idx] {
		r.depth[idx] = z
		// TODO use material of entity
		r.screen.SetPixel(x, y, material.Pixel{Symbol: '~', Color: material.GetColor(85, 133, 212)})
	}
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
		utils.BresenhamWithT(ax, ay, bx, by, func(x, y int, t float64) { r.setPixel(x, y, a.Z*(1-t)+b.Z*t) })
		utils.BresenhamWithT(bx, by, cx, cy, func(x, y int, t float64) { r.setPixel(x, y, b.Z*(1-t)+c.Z*t) })
		utils.BresenhamWithT(cx, cy, ax, ay, func(x, y int, t float64) { r.setPixel(x, y, a.Z*(1-t)+b.Z*t) })

	}
}
