package types

import "time"

type Storage struct {
	// Rate per 1000 ms or time.ParseDuration("1s")
	ResourceRate *Material
	// Current storage container
	ResourceContainer *Material
	// Limit for the resource container
	ResourceContainerLimit *Material
	TimeLastCaptured       time.Time
}

func GenerateEmptyStorage() *Storage {
	return &Storage{
		ResourceContainer:      &Material{},
		ResourceRate:           &Material{},
		ResourceContainerLimit: &Material{},
		TimeLastCaptured:       time.Now(),
	}
}

func (r *Storage) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.TimeLastCaptured)

	resourceDiff := r.ResourceRate.getMaterialDifferenceFromDuration(&durationSinceLastCalculation)
	r.addMaterial(resourceDiff)
}

func (r *Storage) addMaterial(m *Material) {
	r.ResourceContainer.Iron += m.Iron
	r.ResourceContainer.Copper += m.Copper
	r.ResourceContainer.Coal += m.Coal
	r.ResourceContainer.Water += m.Water
}

func (r *Storage) TransferResourceFrom(ext *Storage) {
	timeOnCalculation := time.Now()
	r.calculateOnResourceRate(&timeOnCalculation)
	r.TimeLastCaptured = timeOnCalculation

	r.addMaterial(ext.ResourceContainer)
}
