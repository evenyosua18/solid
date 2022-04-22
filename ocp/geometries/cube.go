package geometries

type Cube struct {
	name string
	side float64
}

func (c Cube) GetName() string {
	return c.name
}

func (c Cube) CalculateVolume() float64 {
	return c.side * c.side * c.side
}

func (c *Cube) GetSide() float64 {
	return c.side
}

func (c *Cube) Create(name string, side float64) {
	c.name = name
	c.side = side
}
