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

func Translate3D(tx, ty, tz float64) Matrix {
	return InitMatrix([][]float64{
		{1, 0, 0, tx},
		{0, 1, 0, ty},
		{0, 0, 1, tz},
		{0, 0, 0, 1},
	})
}

func Scale3D(sx, sy, sz float64) Matrix {
	return InitMatrix([][]float64{
		{sx, 0, 0, 0},
		{0, sy, 0, 0},
		{0, 0, sz, 0},
		{0, 0, 0, 1},
	})
}

func RotateX(a float64) Matrix {
	c := math.Cos(a)
	s := math.Sin(a)
	return InitMatrix([][]float64{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	})
}

func RotateY(a float64) Matrix {
	c := math.Cos(a)
	s := math.Sin(a)
	return InitMatrix([][]float64{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	})
}

func RotateZ(a float64) Matrix {
	c := math.Cos(a)
	s := math.Sin(a)
	return InitMatrix([][]float64{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func (m *Matrix) TransformPoint3D(v Vector3) Vector3 {
	x := v.X
	y := v.Y
	z := v.Z

	// Multiply by matrix (4×4)
	xp := m.M[0][0]*x + m.M[0][1]*y + m.M[0][2]*z + m.M[0][3]
	yp := m.M[1][0]*x + m.M[1][1]*y + m.M[1][2]*z + m.M[1][3]
	zp := m.M[2][0]*x + m.M[2][1]*y + m.M[2][2]*z + m.M[2][3]
	wp := m.M[3][0]*x + m.M[3][1]*y + m.M[3][2]*z + m.M[3][3]

	// Perspective divide
	if wp != 0 {
		xp /= wp
		yp /= wp
		zp /= wp
	}

	return Vector3{X: xp, Y: yp, Z: zp}
}

func Perspective(fov, aspect, near, far float64) Matrix {
	f := 1.0 / math.Tan(fov/2)

	return InitMatrix([][]float64{
		{f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), (2 * far * near) / (near - far)},
		{0, 0, -1, 0},
	})
}

func LookAt(eye, target, up Vector3) Matrix {
	f := target.Sub(eye).Normalize()
	r := Cross(f, up).Normalize()
	u := Cross(r, f)

	return InitMatrix([][]float64{
		{r.X, u.X, -f.X, -Dot(r, eye)},
		{r.Y, u.Y, -f.Y, -Dot(u, eye)},
		{r.Z, u.Z, -f.Z, Dot(f, eye)},
		{0, 0, 0, 1},
	})
}
