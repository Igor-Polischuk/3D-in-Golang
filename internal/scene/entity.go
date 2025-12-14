package scene

import (
	"gortex/internal/geom"
	"gortex/internal/mesh"
)

type Entity struct {
	Mesh mesh.Mesh

	Pos, Rot, Scale geom.Vector3
}

func NewEntity(m mesh.Mesh, pos geom.Vector3) Entity {
	return Entity{
		Mesh:  m,
		Pos:   pos,
		Rot:   geom.ZeroVector3(),
		Scale: geom.GetVector3(1, 1, 1),
	}
}

func (c *Entity) ModelMatrix() geom.Matrix {
	T := geom.Translate3D(c.Pos.X, c.Pos.Y, c.Pos.Z)
	Rx := geom.RotateX(c.Rot.X)
	Ry := geom.RotateY(c.Rot.Y)
	Rz := geom.RotateZ(c.Rot.Z)
	S := geom.Scale3D(c.Scale.X, c.Scale.Y, c.Scale.Z)

	return T.Mul(&Ry).Mul(&Rx).Mul(&Rz).Mul(&S)
}
