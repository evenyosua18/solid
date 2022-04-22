package geometries

type GeometryInterface interface {
	CalculateVolume() float64
	GetName() string
	SetHeight(float64)
	SetRadius(float64)
}
