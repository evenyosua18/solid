package geometries

type Cube struct {
	Geometry
	height float64
}

func (c *Cube) GetName() string {
	return c.name
}

func (c *Cube) SetHeight(height float64) {
	c.height = height
}

func (c *Cube) SetRadius(radius float64) {

}

func (c *Cube) GetHeight() float64 {
	return c.height
}

func (c *Cube) CalculateVolume() float64 {
	return c.height * c.height * c.height
}

func (c *Cube) Create(name string) {
	c.name = name
}
