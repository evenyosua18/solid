package ocp

import (
	"solid/ocp/calculator"
	"solid/ocp/geometries"
)

func Run() {
	//create geometries
	var sphere geometries.Sphere
	var cube geometries.Cube
	var tube geometries.Tube
	var cone geometries.Cone

	//define geometries
	sphere.Create("Sphere", 7)
	cube.Create("Cube", 7)
	tube.Create("Tube", 7, 10)
	cone.Create("Cone", 7, 10)

	//define geometries interface
	var _sphere geometries.GeometryInterface = sphere
	var _cube geometries.GeometryInterface = cube
	var _tube geometries.GeometryInterface = tube
	var _cone geometries.GeometryInterface = cone

	//volume calculator
	var vc calculator.VolumeCalculator
	vc.AddGeomteries(_sphere)
	vc.AddGeomteries(_cube)
	vc.AddGeomteries(_tube)
	vc.AddGeomteries(_cone)

	//print the result
	var vp calculator.VolumePrinter
	//fmt.Println(string(vp.JSON(vc.GetGeometriesVolume())))
	vp.PrintVolume(vc.GetGeometriesVolume())
}
