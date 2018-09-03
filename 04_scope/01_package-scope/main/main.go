package main

import (
	"fmt"

	"github.com/hanchiang/GolangTraining/04_scope/01_package-scope/vis"
)

func main() {
	vis.PrintName()
	fmt.Println("MyName: " + vis.MyName)
	fmt.Println("My place: " + myPlace)
}
