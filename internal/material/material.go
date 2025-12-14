package material

import "gortex/internal/geom"

type FragmentInput struct {
	X, Y int
	Z    float64

	Normal geom.Vector3
}

type Pixel struct {
	Color  Color
	Symbol rune
}
