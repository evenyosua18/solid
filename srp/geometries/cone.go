package geometries

type Cone struct {
	height float64
	radius float64
}

func (c *Cone) GetHeight() float64 {
	return c.height
}

func (c *Cone) GetRadius() float64 {
	return c.radius
}

func (c *Cone) Create(height, radius float64) {
	c.height = height
	c.radius = radius
}
