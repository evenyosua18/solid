package geometries

type Cone struct {
	Geometry
	radius float64
	height float64
}

func (c *Cone) GetRadius() float64 {
	return c.radius
}

func (c *Cone) GetHeight() float64 {
	return c.height
}

func (c *Cone) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Cone) SetHeight(height float64) {
	c.height = height
}

func (c *Cone) Create(name string) {
	c.name = name
}

func (c *Cone) CalculateVolume() float64 {
	return PHI * c.radius * c.radius * c.height / 3
}
