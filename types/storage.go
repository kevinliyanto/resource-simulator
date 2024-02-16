package types

import "time"

type Storage struct {
	Resource *Material
	// Rate per 1000 ms or time.ParseDuration("1s")
	ResourceRate     *Material
	ResourceLimit    *Material
	TimeLastCaptured time.Time
}

func GenerateEmptyResource() *Storage {
	return &Storage{
		Resource:         &Material{},
		ResourceRate:     &Material{},
		ResourceLimit:    &Material{},
		TimeLastCaptured: time.Now(),
	}
}

func (r *Storage) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.TimeLastCaptured)

	resourceDiff := r.ResourceRate.calculateMaterialDiff(&durationSinceLastCalculation)
	r.addMaterial(resourceDiff)
}

func (r *Storage) addMaterial(m *Material) {
	r.Resource.Iron += m.Iron
	r.Resource.Copper += m.Copper
	r.Resource.Coal += m.Coal
	r.Resource.Water += m.Water
}

func (r *Storage) AddResource(ext *Storage) {
	timeOnCalculation := time.Now()

	r.calculateOnResourceRate(&timeOnCalculation)

	r.Resource.Iron += ext.Resource.Iron
	r.Resource.Copper += ext.Resource.Copper
	r.Resource.Coal += ext.Resource.Coal
	r.Resource.Water += ext.Resource.Water

	r.TimeLastCaptured = timeOnCalculation
}
