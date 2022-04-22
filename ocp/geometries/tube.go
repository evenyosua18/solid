package geometries

type Tube struct {
	name   string
	radius float64
	height float64
}

func (t Tube) GetName() string {
	return t.name
}

func (t Tube) CalculateVolume() float64 {
	return PHI * t.radius * t.radius * t.height
}

func (t *Tube) GetRadius() float64 {
	return t.radius
}

func (t *Tube) GetHeight() float64 {
	return t.height
}

func (t *Tube) Create(name string, radius, height float64) {
	t.name = name
	t.radius = radius
	t.height = height
}
