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

func NewTriangle() Mesh {
	return Mesh{
		Vertices: []geom.Vector3{
			{X: 0, Y: 0.5, Z: 0},
			{X: -0.5, Y: -0.5, Z: 0},
			{X: 0.5, Y: -0.5, Z: 0},
		},
		Indices: []int{0, 1, 2},
	}
}

func NewSquare() Mesh {
	return Mesh{
		Vertices: []geom.Vector3{
			{X: -1, Y: 1, Z: 0},  // 0
			{X: 1, Y: 1, Z: 0},   // 1
			{X: 1, Y: -1, Z: 0},  // 2
			{X: -1, Y: -1, Z: 0}, // 3
		},
		Indices: []int{
			0, 1, 2,
			2, 3, 0,
		},
	}
}

func NewCube() Mesh {
	verts := []geom.Vector3{
		{X: -1, Y: -1, Z: -1},
		{X: 1, Y: -1, Z: -1},
		{X: 1, Y: 1, Z: -1},
		{X: -1, Y: 1, Z: -1},
		{X: -1, Y: -1, Z: 1},
		{X: 1, Y: -1, Z: 1},
		{X: 1, Y: 1, Z: 1},
		{X: -1, Y: 1, Z: 1},
	}

	idx := []int{
		// front
		0, 1, 2, 2, 3, 0,
		// back
		4, 5, 6, 6, 7, 4,
		// left
		0, 4, 7, 7, 3, 0,
		// right
		1, 5, 6, 6, 2, 1,
		// top
		3, 2, 6, 6, 7, 3,
		// bottom
		0, 1, 5, 5, 4, 0,
	}

	return Mesh{Vertices: verts, Indices: idx}
}
