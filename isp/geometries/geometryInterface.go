package geometries

type GeometryInterface interface {
	CalculateVolume() float64
	GetName() string
}

type ConeInterface interface {
	GeometryInterface
	SetHeight(float64)
	SetRadius(float64)
}

type CubeInterface interface {
	GeometryInterface
	SetSide(float64)
}

type SphereInterface interface {
	GeometryInterface
	SetRadius(float64)
}

type TubeInterface interface {
	GeometryInterface
	SetRadius(float64)
	SetHeight(float64)
}
