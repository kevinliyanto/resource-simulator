package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	types "github.com/kevinliyanto/resource-simulator/types"
)

func addResourceConcurrently(wg *sync.WaitGroup, originResource *types.Storage, additionalResource *types.Storage) {
	defer wg.Done()

	originResource.AddResource(additionalResource)
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	materials := Generate1000RandomMaterials()

	sumMaterial := &types.Material{}

	for _, v := range materials {
		sumMaterial.Iron += v.Iron
		sumMaterial.Copper += v.Copper
		sumMaterial.Coal += v.Coal
		sumMaterial.Water += v.Water
	}

	fmt.Println("Random material sum", sumMaterial)

	r1 := GenerateRandomResource()
	fmt.Println("r1", r1.Resource)

	r2 := types.GenerateEmptyResource()
	fmt.Println("r2", r2.Resource)

	fmt.Println("Adding resource from r2 to r1 after 10 seconds")

	duration, _ := time.ParseDuration("10s")
	time.Sleep(duration)

	r1.AddResource(r2)

	fmt.Println("r1", r1.Resource, "time on capture", r1.TimeLastCaptured.UnixMilli())

	fmt.Println("=== Adding random materials ===")

	timeBeforeAddRandom := r1.TimeLastCaptured.UnixMilli()
	ironBeforeAddRandom := r1.Resource.Iron
	fmt.Println("r1", r1.Resource, "time on capture", timeBeforeAddRandom)

	for i, v := range materials {
		r1.AddResource(&types.Storage{
			Resource: v,
		})
		time.Sleep(1 * time.Millisecond)

		if i%1000 == 0 {
			fmt.Println("Waiting for resource batch", i/1000, "done")
		}
	}

	timeAfterAddRandom := r1.TimeLastCaptured.UnixMilli()
	ironAfterAddRandom := r1.Resource.Iron
	fmt.Println("r1", r1.Resource, "time on capture", timeAfterAddRandom)

	delta := ironAfterAddRandom - (ironBeforeAddRandom + sumMaterial.Iron)
	deltaRate := delta * 1000 / (float64(timeAfterAddRandom - timeBeforeAddRandom))

	rateDiff := math.Abs(deltaRate - r1.ResourceRate.Iron)
	fmt.Println("Rate during add random", deltaRate)
	fmt.Println("Rate diff due to floating point rounding", rateDiff, "per second")

	fmt.Println("=== Concurrent ===")

	timeBeforeAddRandom = r1.TimeLastCaptured.UnixMilli()
	ironBeforeAddRandom = r1.Resource.Iron
	fmt.Println("r1", r1.Resource, "time on capture", timeBeforeAddRandom)

	fmt.Println("Adding random materials concurrently")

	for i := 0; i < 100; i++ {
		waitGroupSize := 1000

		var wg sync.WaitGroup
		wg.Add(waitGroupSize)

		for materialIdx := 0; materialIdx < waitGroupSize; materialIdx++ {
			createdResource := &types.Storage{
				Resource: materials[i*10+materialIdx],
			}

			go addResourceConcurrently(&wg, r1, createdResource)
		}
		wg.Wait()

		fmt.Println("Waiting for resource batch", i, "done")
	}

	timeAfterAddRandom = r1.TimeLastCaptured.UnixMilli()
	ironAfterAddRandom = r1.Resource.Iron
	fmt.Println("r1", r1.Resource, "time on capture", timeAfterAddRandom)

	delta = ironAfterAddRandom - (ironBeforeAddRandom + sumMaterial.Iron)
	deltaRate = delta * 1000 / (float64(timeAfterAddRandom - timeBeforeAddRandom))

	rateDiff = math.Abs(deltaRate - r1.ResourceRate.Iron)
	fmt.Println("Rate during add random concurrent without proper locking", deltaRate)
	fmt.Println("Rate diff due to floating point rounding & non-proper resource locking", rateDiff, "per second")
}
