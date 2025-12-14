package material

import "gortex/internal/geom"

type FragmentInput struct {
	X, Y int
	Z    float64

	Normal geom.Vector3
}

type Material interface {
	Shade(input FragmentInput) Pixel
}

type Pixel struct {
	Color  Color
	Symbol rune
}

type FillMaterial struct {
	pixel Pixel
}

func NewColorFillMaterial(color Color) FillMaterial {
	return FillMaterial{
		pixel: Pixel{
			Color:  color,
			Symbol: '@',
		},
	}
}

func NewASCIIFillMaterial(s rune) FillMaterial {
	return FillMaterial{
		pixel: Pixel{
			Color:  GetColor(255, 255, 255),
			Symbol: s,
		},
	}
}

func (m FillMaterial) Shade(input FragmentInput) Pixel {
	return m.pixel
}
