package geometries

const (
	PHI = 3.14
)

type GeometryInterface interface {
	CalculateVolume() float64
	GetName() string
}
