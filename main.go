package main

import (
	"fmt"
	"time"
)

func main() {
	r1 := GenerateRandomResource()
	fmt.Println(r1.resource)

	r2 := GenerateEmptyResource()
	fmt.Println(r2.resource)

	fmt.Println("Adding resource from r2 to r1 after 10 seconds")

	duration, _ := time.ParseDuration("10s")
	time.Sleep(duration)

	r1.AddResource(r2)
	fmt.Println(r1.resource)
}
