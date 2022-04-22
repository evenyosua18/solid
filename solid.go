package main

import (
	"fmt"
	"solid/ocp"
	"solid/srp"
)

func main() {
	//SRP
	fmt.Println("Single Responsibility Principle")
	fmt.Println("===============================")
	srp.Run()
	fmt.Println()

	//OCP
	fmt.Println("Open Closed Principle")
	fmt.Println("=====================")
	ocp.Run()
	fmt.Println()
}
