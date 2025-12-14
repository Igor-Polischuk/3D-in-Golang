package mesh

import "gortex/internal/geom"

type Mesh struct {
	Vertices []geom.Vector3
	Indices  []int
}

func (m *Mesh) Transform(MVP geom.Matrix) []geom.Vector3 {
	transformed := make([]geom.Vector3, len(m.Vertices))

	for i, v := range m.Vertices {
		transformed[i] = MVP.TransformPoint3(v)
	}

	return transformed
}
