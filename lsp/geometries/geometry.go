package geometries

const (
	PHI = 3.14
)

type Geometry struct {
	name string
}

func (g Geometry) GetName() string {
	return g.name
}
