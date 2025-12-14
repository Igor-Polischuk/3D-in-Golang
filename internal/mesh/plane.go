package mesh

import "gortex/internal/geom"

func NewSquare(size float64) Mesh {
	h := size / 2

	return Mesh{
		Vertices: []geom.Vector3{
			{X: -h, Y: +h, Z: 0}, // 0 top-left
			{X: +h, Y: +h, Z: 0}, // 1 top-right
			{X: +h, Y: -h, Z: 0}, // 2 bottom-right
			{X: -h, Y: -h, Z: 0}, // 3 bottom-left
		},
		Indices: []int{
			0, 1, 2,
			2, 3, 0,
		},
	}
}
