package main

import (
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

func (r *Resource) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.timeLastCaptured)

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
	// Calculated up to nanosecond for most precise calculation
	duration := float64(d.Nanoseconds()) / 1e9

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

	r.timeLastCaptured = timeOnCalculation
}
