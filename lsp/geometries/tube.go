package geometries

type Tube struct {
	Geometry
	height float64
	radius float64
}

func (t *Tube) SetHeight(height float64) {
	t.height = height
}

func (t *Tube) SetRadius(radius float64) {
	t.radius = radius
}

func (t *Tube) GetHeight() float64 {
	return t.height
}

func (t *Tube) GetRadius() float64 {
	return t.radius
}

func (t *Tube) Create(name string) {
	t.name = name
}

func (t *Tube) CalculateVolume() float64 {
	return PHI * t.radius * t.radius * t.height
}
