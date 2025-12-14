package scene

type Scene struct {
	Entities []*Entity
}

func New() Scene {
	return Scene{Entities: []*Entity{}}
}

func (s *Scene) Add(entities ...*Entity) {
	s.Entities = append(s.Entities, entities...)
}
