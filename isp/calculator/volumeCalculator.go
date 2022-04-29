package calculator

import "solid/isp/geometries"

type VolumeResult map[string]float64

type VolumeCalculator struct {
	geometries []geometries.GeometryInterface
}

func (vc *VolumeCalculator) AddGeometry(result geometries.GeometryInterface) {
	vc.geometries = append(vc.geometries, result)
}

func (vc *VolumeCalculator) GetGeometriesVolume() VolumeResult {
	result := make(VolumeResult)
	for i := 0; i < len(vc.geometries); i++ {
		result[vc.geometries[i].GetName()] = vc.geometries[i].CalculateVolume()
	}

	return result
}
