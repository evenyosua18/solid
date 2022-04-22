package srp

import (
	"solid/srp/calculator"
	"solid/srp/geometries"
)

func Run() {
	//create geometries
	var sphere geometries.Sphere
	var cube geometries.Cube
	var tube geometries.Tube
	var cone geometries.Cone

	sphere.Create(7)
	cube.Create(7)
	tube.Create(10, 7)
	cone.Create(10, 7)

	//calculate geometries
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
