package camera

import "gortex/internal/geom"

type Camera struct {
	Pos   geom.Vector2
	Angle float64
	Zoom  float64
}

func GetCamera(x, y, z float64) Camera {
	return Camera{
		Pos:   geom.Vector2{X: x, Y: y},
		Angle: 0,
		Zoom:  z,
	}
}

// ViewMatrix returns the view matrix.
// This is the INVERSE of camera movements, because the camera doesn't move itself—it moves the entire world in the opposite direction.
func (c *Camera) ViewMatrix() geom.Matrix {
	// 1. Camera model: T * R * S
	t := geom.Translate(c.Pos.X, c.Pos.Y)
	// Translation to world coordinates.
	// If the camera is at point (2,3), this means
	// ALL objects need to be shifted by (-2, -3).

	r := geom.Rotate(c.Angle)
	// Camera rotation. If Angle = 90°, the camera ROTATES,
	// but in the view matrix this causes the opposite rotation of the world.

	s := geom.Scale(c.Zoom, c.Zoom)
	// Scale: enlarges/reduces the entire scene.

	// 2. Combine transformations:
	m := t.Mul(&r).Mul(&s)
	// Order is critical: first Scale → Rotate → Translate
	// but due to left-hand multiplication we write Translate * Rotate * Scale.
	// This gives the correct transformation in WORLD SPACE.

	// 3. View = inverse(ModelCamera)
	inv, _ := m.Inverse()
	// Mathematically:
	// View = (T * R * S)^(-1) = S^-1 * R^-1 * T^-1
	// This is correct: first the world moves in reverse, then rotates in reverse, then scales in reverse.

	return inv
}
