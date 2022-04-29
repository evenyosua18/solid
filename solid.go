package main

import (
	"fmt"
	"solid/isp"
	"solid/lsp"
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

	//LSP
	fmt.Println("Liskov Substitution Principle")
	fmt.Println("=============================")
	lsp.Run()
	fmt.Println()

	//ISP
	fmt.Println("Interface Segregation Principle")
	fmt.Println("===============================")
	isp.Run()
	fmt.Println()
}
