package geometry

type Cube struct {
	side float64
}

func (c *Cube) GetSide() float64 {
	return c.side
}

func (c *Cube) Create(side float64) {
	c.side = side
}
