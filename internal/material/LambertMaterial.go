package material

import "gortex/internal/geom"

type LambertMaterial struct {
	Color    Color
	LightDir geom.Vector3 // у view space
	Ambient  float64
}

func NewLambert(color Color, lightDir geom.Vector3) LambertMaterial {
	return LambertMaterial{
		Color:    color,
		LightDir: lightDir.Normalize(),
		Ambient:  0.2,
	}
}

func (m LambertMaterial) Shade(in FragmentInput) Pixel {
	n := in.Normal.Normalize()
	l := m.LightDir

	diffuse := geom.Dot(n, l)
	if diffuse < 0 {
		diffuse = 0
	}

	intensity := m.Ambient + diffuse
	if intensity > 1 {
		intensity = 1
	}

	color := geom.GetVector3(
		float64(m.Color.R),
		float64(m.Color.G),
		float64(m.Color.B),
	).MulScalar(intensity)

	return Pixel{
		Color:  GetColor(uint8(color.X), uint8(color.Y), uint8(color.Z)),
		Symbol: '█',
	}
}
