package main

import (
	"fmt"
	math_rand "math/rand"
	"time"
)

// type Iron float64
// type Copper float64
// type Coal float64
// type Water float64

type Material struct {
	Iron   float64
	Copper float64
	Coal   float64
	Water  float64
}

type Resource struct {
	resource *Material
	// Rate per 1000 ms or time.ParseDuration("1s")
	resourceRate     *Material
	resourceLimit    *Material
	timeLastCaptured time.Time
}

func GenerateEmptyResource() *Resource {
	return &Resource{
		resource:         &Material{},
		resourceRate:     &Material{},
		resourceLimit:    &Material{},
		timeLastCaptured: time.Now(),
	}
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

func (r *Resource) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.timeLastCaptured)

	fmt.Println("Duration since last calculation", durationSinceLastCalculation)

	resourceDiff := r.resourceRate.calculateMaterialDiff(&durationSinceLastCalculation)
	r.addMaterial(resourceDiff)
}

func (r *Resource) addMaterial(m *Material) {
	r.resource.Iron += m.Iron
	r.resource.Copper += m.Copper
	r.resource.Coal += m.Coal
	r.resource.Water += m.Water
}

func (resourceRate *Material) calculateMaterialDiff(d *time.Duration) *Material {
	duration := float64(d.Milliseconds()) / 1000

	return &Material{
		Iron:   resourceRate.Iron * duration,
		Copper: resourceRate.Copper * duration,
		Coal:   resourceRate.Coal * duration,
		Water:  resourceRate.Water * duration,
	}
}

func (r *Resource) AddResource(ext *Resource) {
	timeOnCalculation := time.Now()

	r.calculateOnResourceRate(&timeOnCalculation)

	r.resource.Iron += ext.resource.Iron
	r.resource.Copper += ext.resource.Copper
	r.resource.Coal += ext.resource.Coal
	r.resource.Water += ext.resource.Water
}
