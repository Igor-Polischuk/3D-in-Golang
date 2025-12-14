package render

import (
	"gortex/internal/geom"
	"gortex/internal/material"
	"gortex/internal/mesh"
	"gortex/internal/scene"
	"gortex/internal/utils"
	"math"
)

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

		r.RenderMesh(entity.Mesh, entity.Material, MVP)
	}
}

func ToScreen(v geom.Vector3, w, h int) (int, int) {
	x := (v.X + 1) * 0.5 * float64(w-1)
	y := (v.Y + 1) * 0.5 * float64(h-1)

	return int(x), int(y)
}

func (r Renderer) setPixel(x, y int, z float64, pixel material.Pixel) {
	// Bounds checking to prevent index out of range
	if x < 0 || x >= r.screen.Width() || y < 0 || y >= r.screen.Height() {
		return
	}
	idx := x + y*r.screen.Width()
	if z < r.depth[idx] {
		r.depth[idx] = z
		r.screen.SetPixel(x, y, pixel)
	}
}

func (r Renderer) RenderMesh(mesh mesh.Mesh, mat material.Material, MVP geom.Matrix) {
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
		utils.BresenhamWithT(ax, ay, bx, by, func(x, y int, t float64) {
			z := a.Z*(1-t) + b.Z*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})
		utils.BresenhamWithT(bx, by, cx, cy, func(x, y int, t float64) {
			z := b.Z*(1-t) + c.Z*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})
		utils.BresenhamWithT(cx, cy, ax, ay, func(x, y int, t float64) {
			z := c.Z*(1-t) + a.Z*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})

	}
}
