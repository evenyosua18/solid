package isp

import (
	"solid/isp/calculator"
	"solid/isp/geometries"
)

func Run() {
	//create cube
	var cube geometries.Cube
	cube.Create("cube number one")

	var _cube geometries.CubeInterface = &cube
	_cube.SetSide(7)

	//create cone
	var cone geometries.Cone
	cone.Create("cone number one")

	var _cone geometries.ConeInterface = &cone
	_cone.SetHeight(7)
	_cone.SetRadius(7)

	//create sphere
	var sphere geometries.Sphere
	sphere.Create("sphere number one")

	var _sphere geometries.SphereInterface = &sphere
	_sphere.SetRadius(7)

	//create tube
	var tube geometries.Tube
	tube.Create("tube number one")

	var _tube geometries.TubeInterface = &tube
	_tube.SetRadius(7)
	_tube.SetHeight(7)

	//set calculator
	var volumeCalculator calculator.VolumeCalculator
	volumeCalculator.AddGeometry(_cone)
	volumeCalculator.AddGeometry(_cube)
	volumeCalculator.AddGeometry(_tube)
	volumeCalculator.AddGeometry(_sphere)

	//print result
	var volumePrinter calculator.VolumePrinter
	volumePrinter.Print(volumeCalculator.GetGeometriesVolume())
}
