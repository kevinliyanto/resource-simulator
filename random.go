package main

import (
	math_rand "math/rand"
	"time"

	types "github.com/kevinliyanto/resource-simulator/types"
)

func generateRandomMaterial() *types.Material {
	return &types.Material{
		Iron:   float64(math_rand.Intn(10)),
		Copper: float64(math_rand.Intn(10)),
		Coal:   float64(math_rand.Intn(10)),
		Water:  float64(math_rand.Intn(10)),
	}
}

func Generate1000RandomMaterials() *[100000]*types.Material {
	var arr [100000]*types.Material

	for i := 0; i < 100000; i++ {
		arr[i] = generateRandomMaterial()
	}

	return &arr
}

func GenerateRandomResource() *types.Storage {
	defaultRate := types.Material{
		Iron:   10.0,
		Copper: 10.0,
		Coal:   10.0,
		Water:  5.0,
	}

	defaultLimit := types.Material{
		Iron:   80000.0,
		Copper: 80000.0,
		Coal:   80000.0,
		Water:  24000.0,
	}

	return &types.Storage{
		ResourceContainer: &types.Material{
			Iron:   defaultLimit.Iron * math_rand.Float64(),
			Copper: defaultLimit.Copper * math_rand.Float64(),
			Coal:   defaultLimit.Coal * math_rand.Float64(),
			Water:  defaultLimit.Water * math_rand.Float64(),
		},
		ResourceRate:           &defaultRate,
		ResourceContainerLimit: &defaultLimit,
		TimeLastCaptured:       time.Now(),
	}
}
