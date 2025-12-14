package mesh

import "gortex/internal/geom"

func NewCube(dim geom.Vector3) Mesh {

	rx, ry, rz := dim.X/2, dim.Y/2, dim.Z/2

	m := Mesh{
		Vertices: []geom.Vector3{
			{X: -rx, Y: -ry, Z: -rz}, // 0
			{X: +rx, Y: -ry, Z: -rz}, // 1
			{X: +rx, Y: +ry, Z: -rz}, // 2
			{X: -rx, Y: +ry, Z: -rz}, // 3
			{X: -rx, Y: -ry, Z: +rz}, // 4
			{X: +rx, Y: -ry, Z: +rz}, // 5
			{X: +rx, Y: +ry, Z: +rz}, // 6
			{X: -rx, Y: +ry, Z: +rz}, // 7
		},
		Indices: []int{
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
		},
	}

	return m
}
