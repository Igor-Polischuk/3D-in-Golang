package geom

type Vector2 struct {
	X, Y float64
}

func (v *Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v *Vector2) Sub(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}
