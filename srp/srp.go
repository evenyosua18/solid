package srp

import (
	"solid/srp/calculator"
	"solid/srp/geometry"
)

func Run() {
	//create geometry
	var sphere geometry.Sphere
	var cube geometry.Cube
	var tube geometry.Tube
	var cone geometry.Cone

	sphere.Create(7)
	cube.Create(7)
	tube.Create(7, 7)
	cone.Create(7, 7)

	//calculate geometry
	var volumeCalculator calculator.VolumeCalculator
	var volumeResult calculator.VolumeResult

	volumeCalculator.SetSphere(sphere)
	volumeCalculator.SetCube(cube)
	volumeCalculator.SetTube(tube)
	volumeCalculator.SetCone(cone)

	volumeCalculator.CalculateSphere(&volumeResult)
	volumeCalculator.CalculateCube(&volumeResult)
	volumeCalculator.CalculateTube(&volumeResult)
	volumeCalculator.CalculateCone(&volumeResult)

	//print the result
	var printVolume calculator.VolumePrinter

	printVolume.PrintVolume(volumeResult, calculator.SPHERE)
	printVolume.PrintVolume(volumeResult, calculator.CUBE)
	printVolume.PrintVolume(volumeResult, calculator.TUBE)
	printVolume.PrintVolume(volumeResult, calculator.CONE)
}
