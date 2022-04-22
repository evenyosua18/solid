package geometries

type Cone struct {
	name   string
	height float64
	radius float64
}

func (c Cone) GetName() string {
	return c.name
}

func (c Cone) CalculateVolume() float64 {
	return PHI * c.radius * c.radius * c.height / 3
}

func (c *Cone) GetHeight() float64 {
	return c.height
}

func (c *Cone) GetRadius() float64 {
	return c.radius
}

func (c *Cone) Create(name string, height, radius float64) {
	c.name = name
	c.height = height
	c.radius = radius
}
