package calculator

import (
	"math"
	"solid/srp/geometries"
)

const (
	CONE   = "CONE"
	SPHERE = "SPHERE"
	TUBE   = "TUBE"
	CUBE   = "CUBE"

	PHI = 3.14
)

type VolumeResult struct {
	sphereVolume float64
	coneVolume   float64
	cubeVolume   float64
	tubeVolume   float64
}

type VolumeCalculator struct {
	sphere geometries.Sphere
	cone   geometries.Cone
	cube   geometries.Cube
	tube   geometries.Tube
}

func (v *VolumeCalculator) SetSphere(sphere geometries.Sphere) {
	v.sphere = sphere
}

func (v *VolumeCalculator) SetCube(cube geometries.Cube) {
	v.cube = cube
}

func (v *VolumeCalculator) SetTube(tube geometries.Tube) {
	v.tube = tube
}

func (v *VolumeCalculator) SetCone(cone geometries.Cone) {
	v.cone = cone
}

func (v *VolumeCalculator) CalculateSphere(vr *VolumeResult) {
	vr.sphereVolume = 4 / 3 * PHI * math.Pow(v.sphere.GetRadius(), 3)
}

func (v *VolumeCalculator) CalculateCube(vr *VolumeResult) {
	vr.cubeVolume = math.Pow(v.cube.GetSide(), 3)
}

func (v *VolumeCalculator) CalculateCone(vr *VolumeResult) {
	vr.coneVolume = PHI * v.cone.GetRadius() * v.cone.GetRadius() * v.cone.GetHeight() / 3
}

func (v *VolumeCalculator) CalculateTube(vr *VolumeResult) {
	vr.tubeVolume = PHI * v.tube.GetRadius() * v.tube.GetRadius() * v.tube.GetHeight()
}
