package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	materials := Generate1000RandomMaterials()

	sumMaterial := &Material{}

	for _, v := range materials {
		sumMaterial.Iron += v.Iron
		sumMaterial.Copper += v.Copper
		sumMaterial.Coal += v.Coal
		sumMaterial.Water += v.Water
	}

	fmt.Println("Random material sum", sumMaterial)

	r1 := GenerateRandomResource()
	fmt.Println("r1", r1.resource)

	r2 := GenerateEmptyResource()
	fmt.Println("r2", r2.resource)

	fmt.Println("Adding resource from r2 to r1 after 10 seconds")

	duration, _ := time.ParseDuration("10s")
	time.Sleep(duration)

	r1.AddResource(r2)

	timeBeforeAddRandom := r1.timeLastCaptured.UnixMilli()
	ironBeforeAddRandom := r1.resource.Iron
	fmt.Println("r1", r1.resource, "time on capture", timeBeforeAddRandom)

	fmt.Println("Adding random materials")

	for _, v := range materials {
		r1.AddResource(&Resource{
			resource: v,
		})
		time.Sleep(10 * time.Millisecond)
	}

	timeAfterAddRandom := r1.timeLastCaptured.UnixMilli()
	ironAfterAddRandom := r1.resource.Iron
	fmt.Println("r1", r1.resource, "time on capture", timeAfterAddRandom)

	delta := ironAfterAddRandom - (ironBeforeAddRandom + sumMaterial.Iron)
	deltaRate := delta * 1000 / (float64(timeAfterAddRandom - timeBeforeAddRandom))

	rateDiff := math.Abs(deltaRate - r1.resourceRate.Iron)
	fmt.Println("Rate during add random", deltaRate)
	fmt.Println("Rate diff due to floating point rounding", rateDiff, "per second")
}
