package geometries

type Cube struct {
	Geometry
	side float64
}

func (c *Cube) GetSide() float64 {
	return c.side
}

func (c *Cube) SetSide(side float64) {
	c.side = side
}

func (c *Cube) Create(name string) {
	c.name = name
}

func (c Cube) CalculateVolume() float64 {
	return c.side * c.side * c.side
}
