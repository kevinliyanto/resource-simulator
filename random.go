package main

import (
	math_rand "math/rand"
	"time"
)

func generateRandomMaterial() *Material {
	return &Material{
		Iron:   float64(math_rand.Intn(10)),
		Copper: float64(math_rand.Intn(10)),
		Coal:   float64(math_rand.Intn(10)),
		Water:  float64(math_rand.Intn(10)),
	}
}

func Generate1000RandomMaterials() *[100000]*Material {
	var arr [100000]*Material

	for i := 0; i < 100000; i++ {
		arr[i] = generateRandomMaterial()
	}

	return &arr
}

func GenerateRandomResource() *Resource {
	defaultRate := Material{
		Iron:   10.0,
		Copper: 10.0,
		Coal:   10.0,
		Water:  5.0,
	}

	defaultLimit := Material{
		Iron:   80000.0,
		Copper: 80000.0,
		Coal:   80000.0,
		Water:  24000.0,
	}

	return &Resource{
		resource: &Material{
			Iron:   defaultLimit.Iron * math_rand.Float64(),
			Copper: defaultLimit.Copper * math_rand.Float64(),
			Coal:   defaultLimit.Coal * math_rand.Float64(),
			Water:  defaultLimit.Water * math_rand.Float64(),
		},
		resourceRate:     &defaultRate,
		resourceLimit:    &defaultLimit,
		timeLastCaptured: time.Now(),
	}
}
