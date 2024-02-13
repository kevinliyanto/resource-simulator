package main

import (
	"fmt"
)

func main() {
	r1 := GenerateRandomResource()
	fmt.Println(r1.resource)

	r2 := GenerateRandomResource()
	fmt.Println(r2.resource)

	fmt.Println("Adding resource from r2 to r1")

	r1.AddMaterial(r2.resource)
	fmt.Println(r1.resource)
}
