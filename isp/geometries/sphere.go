package geometries

import "math"

type Sphere struct {
	Geometry
	radius float64
}

func (s *Sphere) SetRadius(radius float64) {
	s.radius = radius
}

func (s *Sphere) GetRadius() float64 {
	return s.radius
}

func (s *Sphere) Create(name string) {
	s.name = name
}

func (s Sphere) CalculateVolume() float64 {
	return 4 / 3 * PHI * math.Pow(s.radius, 3)
}
