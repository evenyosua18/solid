package geometry

type Sphere struct {
	radius float64
}

func (s *Sphere) GetRadius() float64 {
	return s.radius
}

func (s *Sphere) Create(radius float64) {
	s.radius = radius
}
