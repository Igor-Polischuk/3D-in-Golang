package mesh

import "gortex/internal/geom"

type Mesh struct {
	Vertices []geom.Vector3
	Indices  []int
}

type VertexOut struct {
	NDC     geom.Vector3 // координати після perspective divide (X,Y,Z)
	ViewZ   float64      // Z у view space
	ViewPos geom.Vector3
}

func (m Mesh) Transform(model, view, proj geom.Matrix) []VertexOut {
	out := make([]VertexOut, len(m.Vertices))

	for i, v := range m.Vertices {

		// 1) World space
		vw := model.TransformPoint3(v)

		// 2) View space (важливо!)
		vv := view.TransformPoint3(vw)

		// save Z before projection
		viewZ := vv.Z

		// 3) Clip space (projection)
		vc := proj.TransformPoint3(vv)

		// 4) Perspective divide → NDC
		ndc := geom.Vector3{
			X: vc.X,
			Y: vc.Y,
			Z: vc.Z,
		}

		out[i] = VertexOut{
			NDC:     ndc,
			ViewZ:   viewZ,
			ViewPos: vv,
		}
	}

	return out
}
