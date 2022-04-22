package geometries

import "math"

type Sphere struct {
	name   string
	radius float64
}

func (s Sphere) GetName() string {
	return s.name
}

func (s Sphere) CalculateVolume() float64 {
	return 4 / 3 * PHI * math.Pow(s.radius, 3)
}

func (s *Sphere) GetRadius() float64 {
	return s.radius
}

func (s *Sphere) Create(name string, radius float64) {
	s.radius = radius
	s.name = name
}
