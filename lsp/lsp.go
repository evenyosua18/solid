package lsp

import (
	"solid/lsp/calculator"
	"solid/lsp/geometries"
)

func Run() {
	//create cube
	var cube geometries.Cube
	cube.Create("Cube Number One")

	var _cube geometries.GeometryInterface = &cube
	_cube.SetHeight(7)

	var cube2 geometries.Cube
	cube2.Create("Cube Number Two")

	var _cube2 geometries.GeometryInterface = &cube2
	_cube2.SetHeight(14)

	//create cone
	var cone geometries.Cone
	cone.Create("Cone Number One")

	var _cone geometries.GeometryInterface = &cone
	_cone.SetHeight(7)
	_cone.SetRadius(7)

	//create sphere
	var sphere geometries.Sphere
	sphere.Create("Sphere Number One")

	var _sphere geometries.GeometryInterface = &sphere
	_sphere.SetRadius(7)

	//create tube
	var tube geometries.Tube
	tube.Create("Tube Number One")

	var _tube geometries.GeometryInterface = &tube
	_tube.SetRadius(7)
	_tube.SetHeight(7)

	//calculator
	var volumeCalculator calculator.VolumeCalculator

	//add geometries
	volumeCalculator.AddGeometry(_cube)
	volumeCalculator.AddGeometry(_cube2)
	volumeCalculator.AddGeometry(_cone)
	volumeCalculator.AddGeometry(_sphere)
	volumeCalculator.AddGeometry(_tube)

	//print
	var printer calculator.VolumePrinter
	printer.Print(volumeCalculator.GetGeometriesVolume())
}
