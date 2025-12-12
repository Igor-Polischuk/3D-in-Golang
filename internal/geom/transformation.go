package geom

import "math"

func Translate(tx, ty float64) Matrix {
	return InitMatrix([][]float64{
		{1, 0, tx},
		{0, 1, ty},
		{0, 0, 1},
	})
}

func Rotate(a float64) Matrix {
	c := math.Cos(a)
	s := math.Sin(a)
	return InitMatrix([][]float64{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	})
}

func Scale(sx, sy float64) Matrix {
	return InitMatrix([][]float64{
		{sx, 0, 0},
		{0, sy, 0},
		{0, 0, 1},
	})
}

func (m *Matrix) TransformPoint(v Vector2) Vector2 {
	x := v.X
	y := v.Y
	nx := m.M[0][0]*x + m.M[0][1]*y + m.M[0][2]
	ny := m.M[1][0]*x + m.M[1][1]*y + m.M[1][2]
	return Vector2{X: nx, Y: ny}
}
