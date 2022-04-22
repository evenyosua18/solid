package calculator

import "solid/ocp/geometries"

type VolumeResult map[string]float64

type VolumeCalculator struct {
	geometries []geometries.GeometryInterface
}

func (vc *VolumeCalculator) AddGeomteries(g geometries.GeometryInterface) {
	vc.geometries = append(vc.geometries, g)
}

func (vc *VolumeCalculator) GetGeometriesVolume() VolumeResult {
	result := make(VolumeResult)

	for i := 0; i < len(vc.geometries); i++ {
		result[vc.geometries[i].GetName()] = vc.geometries[i].CalculateVolume()
	}

	return result
}
