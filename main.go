package main

import (
	"fmt"
	"sync"
	"time"

	types "github.com/kevinliyanto/resource-simulator/types"
)

func addResourceConcurrently(wg *sync.WaitGroup, originResource *types.Storage, additionalResource *types.Storage) {
	defer wg.Done()

	originResource.TransferResourceFrom(additionalResource)
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
	fmt.Println("r1", r1.ResourceContainer)

	r2 := types.GenerateEmptyStorage()
	fmt.Println("r2", r2.ResourceContainer)

	fmt.Println("Adding resource from r2 to r1 after 10 seconds")

	time.Sleep(10 * time.Second)

	r1.TransferResourceFrom(r2)

	fmt.Println("r1", r1.ResourceContainer, "time on capture", r1.TimeLastCaptured.UnixMilli())

	fmt.Println("=== Adding random materials ===")

	r1StorageStateInitial := getStorageState(r1)
	fmt.Println("r1 storage state", r1StorageStateInitial.printStatus())

	for i, v := range materials {
		r1.TransferResourceFrom(&types.Storage{
			ResourceContainer: v,
		})
		time.Sleep(1 * time.Millisecond)

		if i%1000 == 0 {
			fmt.Println("Waiting for resource batch", i/1000, "done")
		}
	}

	r1StorageStateFinal := getStorageState(r1)
	fmt.Println("r1 storage state", r1StorageStateFinal.printStatus())

	deltaStorageState := r1StorageStateFinal.deltaStorageState(r1StorageStateInitial, sumMaterial)

	fmt.Println("r1 storage rate differential", deltaStorageState.printStatus())

	fmt.Println("=== Concurrent ===")

	r1StorageStateInitial = getStorageState(r1)
	fmt.Println("r1 storage state", r1StorageStateInitial.printStatus())

	fmt.Println("Adding random materials concurrently")

	for i := 0; i < 100; i++ {
		waitGroupSize := 1000

		var wg sync.WaitGroup
		wg.Add(waitGroupSize)

		for materialIdx := 0; materialIdx < waitGroupSize; materialIdx++ {
			createdResource := &types.Storage{
				ResourceContainer: materials[i*10+materialIdx],
			}

			go addResourceConcurrently(&wg, r1, createdResource)
		}
		wg.Wait()

		fmt.Println("Waiting for resource batch", i, "done")
	}

	r1StorageStateFinal = getStorageState(r1)
	fmt.Println("r1 storage state", r1StorageStateFinal.printStatus())

	deltaStorageState = r1StorageStateFinal.deltaStorageState(r1StorageStateInitial, sumMaterial)

	fmt.Println("r1 storage rate differential", deltaStorageState.printStatus())
}
