package main

import "time"

type Storage struct {
	resource *Material
	// Rate per 1000 ms or time.ParseDuration("1s")
	resourceRate     *Material
	resourceLimit    *Material
	timeLastCaptured time.Time
}

func GenerateEmptyResource() *Storage {
	return &Storage{
		resource:         &Material{},
		resourceRate:     &Material{},
		resourceLimit:    &Material{},
		timeLastCaptured: time.Now(),
	}
}

func (r *Storage) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.timeLastCaptured)

	resourceDiff := r.resourceRate.calculateMaterialDiff(&durationSinceLastCalculation)
	r.addMaterial(resourceDiff)
}

func (r *Storage) addMaterial(m *Material) {
	r.resource.Iron += m.Iron
	r.resource.Copper += m.Copper
	r.resource.Coal += m.Coal
	r.resource.Water += m.Water
}

func (r *Storage) AddResource(ext *Storage) {
	timeOnCalculation := time.Now()

	r.calculateOnResourceRate(&timeOnCalculation)

	r.resource.Iron += ext.resource.Iron
	r.resource.Copper += ext.resource.Copper
	r.resource.Coal += ext.resource.Coal
	r.resource.Water += ext.resource.Water

	r.timeLastCaptured = timeOnCalculation
}
