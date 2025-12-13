package geom

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func (v *Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v *Vector3) Sub(other Vector3) Vector3 {
	return Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v Vector3) MulScalar(s float64) Vector3 {
	return Vector3{v.X * s, v.Y * s, v.Z * s}
}

func Dot(a, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a, b Vector3) Vector3 {
	return Vector3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3) Normalize() Vector3 {
	l := v.Length()
	if l == 0 {
		return v
	}
	return Vector3{v.X / l, v.Y / l, v.Z / l}
}
