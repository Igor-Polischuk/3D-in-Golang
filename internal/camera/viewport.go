package camera

import "gortex/internal/geom"

// ViewportMatrix compensates for the fact that terminal pixels are not square.
// aspect = width/height
// pixelAspect = actual aspect ratio of the character (height/width).
func ViewportMatrix(aspect, pixelAspect float64) geom.Matrix {
	// X needs to be "compressed" or "stretched" so that a square in world space remains a square.
	// If we DON'T apply this matrix, all objects will be stretched.

	// Scaling along X:
	// scaleX = aspect * pixelAspect
	//
	// Why?
	// pixelAspect ≈ 11/24 (ASCII characters are tall and narrow)
	// aspect — width/height ratio in NDC
	//
	// Together they form a correction factor.

	scaleX := aspect * pixelAspect

	return geom.Scale(scaleX, 1)
	// This makes the X coordinate "equal" to Y in geometric terms.
}
