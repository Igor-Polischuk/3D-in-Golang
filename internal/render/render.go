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

func (r Renderer) RenderScene(scn scene.Scene, v, p geom.Matrix) {
	for i := range r.depth {
		r.depth[i] = math.Inf(1) // +∞
	}
	for _, entity := range scn.Entities {
		model := entity.ModelMatrix()
		// MVP := vp.Mul(&model)

		r.Render(entity.Mesh, entity.Material, model, v, p)
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
	d := -z
	if d < r.depth[idx] {
		r.depth[idx] = d
		r.screen.SetPixel(x, y, pixel)
	}
}

func edge(ax, ay, bx, by, px, py int) int {
	return (px-ax)*(by-ay) - (py-ay)*(bx-ax)
}

func (r Renderer) fillTriangle(
	a, b, c mesh.VertexOut,
	mat material.Material,
) {
	ab := b.ViewPos.Sub(a.ViewPos)
	ac := c.ViewPos.Sub(a.ViewPos)
	normal := geom.Cross(ab, ac).Normalize()

	// if normal.Z >= 0 {
	// 	return
	// }

	ax, ay := ToScreen(a.NDC, r.screen.Width(), r.screen.Height())
	bx, by := ToScreen(b.NDC, r.screen.Width(), r.screen.Height())
	cx, cy := ToScreen(c.NDC, r.screen.Width(), r.screen.Height())

	minX := max(0, min(ax, bx, cx))
	maxX := min(r.screen.Width()-1, max(ax, bx, cx))
	minY := max(0, min(ay, by, cy))
	maxY := min(r.screen.Height()-1, max(ay, by, cy))

	area := edge(ax, ay, bx, by, cx, cy)
	if area == 0 {
		return
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {

			w0 := edge(bx, by, cx, cy, x, y)
			w1 := edge(cx, cy, ax, ay, x, y)
			w2 := edge(ax, ay, bx, by, x, y)

			if (w0 >= 0 && w1 >= 0 && w2 >= 0) ||
				(w0 <= 0 && w1 <= 0 && w2 <= 0) {

				wa := float64(w0) / float64(area)
				wb := float64(w1) / float64(area)
				wc := float64(w2) / float64(area)

				z := wa*a.ViewZ + wb*b.ViewZ + wc*c.ViewZ
				pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z, Normal: normal})
				r.setPixel(x, y, z, pixel)
			}
		}
	}
}

func (r Renderer) Render(mesh mesh.Mesh, mat material.Material, model, view, proj geom.Matrix) {
	transformed := mesh.Transform(model, view, proj)

	for i := 0; i < len(mesh.Indices); i += 3 {
		a := transformed[mesh.Indices[i]]
		b := transformed[mesh.Indices[i+1]]
		c := transformed[mesh.Indices[i+2]]

		r.fillTriangle(a, b, c, mat)
	}
}

func (r Renderer) RenderMesh(mesh mesh.Mesh, mat material.Material, model, view, proj geom.Matrix) {
	// MVP := proj.Mul(&view).Mul(&model)
	transformed := mesh.Transform(model, view, proj)

	for i := 0; i < len(mesh.Indices); i += 3 {
		a := transformed[mesh.Indices[i]]
		b := transformed[mesh.Indices[i+1]]
		c := transformed[mesh.Indices[i+2]]

		// project to screen
		ax, ay := ToScreen(a.NDC, r.screen.Width(), r.screen.Height())
		bx, by := ToScreen(b.NDC, r.screen.Width(), r.screen.Height())
		cx, cy := ToScreen(c.NDC, r.screen.Width(), r.screen.Height())

		// draw wireframe
		utils.BresenhamWithT(ax, ay, bx, by, func(x, y int, t float64) {
			z := a.ViewZ*(1-t) + b.ViewZ*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})
		utils.BresenhamWithT(bx, by, cx, cy, func(x, y int, t float64) {
			z := b.ViewZ*(1-t) + c.ViewZ*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})
		utils.BresenhamWithT(cx, cy, ax, ay, func(x, y int, t float64) {
			z := c.ViewZ*(1-t) + a.ViewZ*t
			pixel := mat.Shade(material.FragmentInput{X: x, Y: y, Z: z})
			r.setPixel(x, y, z, pixel)
		})

	}
}
