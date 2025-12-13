package render

import (
	"gortex/internal/geom"
	"gortex/internal/mesh"
)

type Screen interface {
	DrawLine(x0, y0, x1, y1 int)
	Width() int
	Height() int
}

func ToScreen(v geom.Vector3, w, h int) (int, int) {
	x := (v.X + 1) * 0.5 * float64(w-1)
	y := (v.Y + 1) * 0.5 * float64(h-1)

	return int(x), int(y)
}

func RenderMesh(mesh mesh.Mesh, MVP geom.Matrix, screen Screen) {
	// MVP := proj.Mul(&view).Mul(&model)
	transformed := mesh.Transform(MVP)

	for i := 0; i < len(mesh.Indices); i += 3 {
		a := transformed[mesh.Indices[i]]
		b := transformed[mesh.Indices[i+1]]
		c := transformed[mesh.Indices[i+2]]

		// project to screen
		ax, ay := ToScreen(a, screen.Width(), screen.Height())
		bx, by := ToScreen(b, screen.Width(), screen.Height())
		cx, cy := ToScreen(c, screen.Width(), screen.Height())

		// draw wireframe
		screen.DrawLine(ax, ay, bx, by)
		screen.DrawLine(bx, by, cx, cy)
		screen.DrawLine(cx, cy, ax, ay)
	}
}
