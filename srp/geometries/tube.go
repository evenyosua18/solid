package geometries

type Tube struct {
	height float64
	radius float64
}

func (t *Tube) GetHeight() float64 {
	return t.height
}

func (t *Tube) GetRadius() float64 {
	return t.radius
}

func (t *Tube) Create(height, radius float64) {
	t.height = height
	t.radius = radius
}
